package service

import (
	"github.com/liyuliang/queue-services"
	"fullSiteSpider/storage"
	"github.com/liyuliang/utils/regex"
	"github.com/liyuliang/dom-parser"
	"github.com/pkg/errors"
	"fullSiteSpider/config"
	"log"
	"time"
	"strings"
	"github.com/liyuliang/utils/request"
	"net/smtp"
	"os"
)

const WaitCrawlerQueue = "WAIT_CRAWLER_QUEUE"
const WaitEmailQueue = "WAIT_SEND_EMAIL_QUEUE"
const EmailSentList = "EMAIL_SEND_LIST"

var EmptyQueueErr = errors.New("Empty queue")
var ThirdPartyErr = errors.New("third party url")

func Init() {

	initTasks()

	services.AddMultiProcessTask("spider run...", func(workerNum int) (err error) {

		uri := popQueue(WaitCrawlerQueue)
		if uri != "" {
			spiderRun(uri)
			return

		} else {
			log.Println(EmptyQueueErr.Error())
			time.Sleep(time.Hour * 10)
			return EmptyQueueErr
		}
	})

	services.AddSingleProcessTask("send email...", func(workerNum int) (err error) {

		mailTo, title, uri := popEmailSendQueue(WaitEmailQueue)
		if mailTo != "" && title != "" && uri != "" {
			content := toMailContent(title, uri)

			err = sendEmail(mailTo, content)

			if err == nil {
				recordEmailSent(uri, title)
			}
		}

		return err
	})
}

func recordEmailSent(uri string, title string) {
	r := storage.Redis()
	r.HSet(EmailSentList, uri, title)
}

func sendEmail(mailTo string, body string) error {
	to := []string{mailTo}

	from := config.Mail().Account
	password := config.Mail().Password
	smtpPort := config.Mail().Port
	smtpHost := config.Mail().SmtpHost

	auth := smtp.PlainAuth("", from, password, smtpHost)

	user := from

	nickname := "发送人名称"
	subject := "邮件通知"
	contentType := "Content-Type: text/plain; charset=UTF-8"

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	err := smtp.SendMail(smtpHost+smtpPort, auth, user, to, msg)
	return err
}

func toMailContent(title string, uri string) string {
	return "<a href=\"" + uri + "\">" + title + "</a>"
}

func initTasks() {
	for _, task := range config.Tasks() {
		addQueue(WaitCrawlerQueue, task.Url)
	}
}

func getTaskByUrl(uri string) (task config.Task) {
	for _, t := range config.Tasks() {

		if strings.Contains(strings.ToLower(uri), strings.ToLower(t.Domain())) {
			task = t
			break
		}
	}
	return task
}

func spiderRun(uri string) (err error) {

	task := getTaskByUrl(uri)

	if isThirdPartyUrl(uri) {
		return ThirdPartyErr
	}

	dom, err := parseHtml(uri)

	//defer func() {
	//	dom.Close()
	//}()

	if err != nil {
		return err
	}

	title := getTitle(dom, task)
	hrefs := getHrefs(dom, task)

	log.Println(title)
	log.Println(hrefs)
	os.Exit(-1)
	r := storage.Redis()

	titleRecord, _ := r.Get(uri)

	if titleRecord == "" {

		checkAndSendEmail(uri, title)

		for _, href := range hrefs {
			href = formatUrl(href)
			addQueue(WaitCrawlerQueue, href)
		}
	}
	return
}
func getHrefs(dom *parser.Dom, task config.Task) (hrefs []string) {

	for _, a := range dom.FindAll(task.HrefsSelector) {
		h, ok := a.Attr("href")
		if ok {
			h = addHost(h, task.TitleSelector)

			hrefs = append(hrefs, h)
		}
	}
	return hrefs
}

func addHost(uri string, host string) string {
	uri = formatUrl(uri)
	if !strings.Contains(uri, "https://") && !strings.Contains(uri, "http://") {
		uri = host + uri
	}
	return uri
}

func getTitle(dom *parser.Dom, task config.Task) string {

	return dom.Find(task.TitleSelector).Text()
}

func parseHtml(uri string) (*parser.Dom, error) {
	resp := request.HttpGet(uri)
	if resp.Err != nil {
		return nil, resp.Err
	}

	return parser.InitDom(resp.Data)
}

func isThirdPartyUrl(uri string) bool {

	isThirdParty := true
	for _, t := range config.Tasks() {
		if strings.Contains(strings.ToLower(uri), strings.ToLower(t.Domain())) {
			isThirdParty = false
			break
		}
	}
	return isThirdParty
}

func checkAndSendEmail(uri string, title string) {
	if hasEmailSent(EmailSentList, uri) {
		return
	}

	for name, task := range config.Tasks() {
		domain := strings.Replace(name, "task.", "", -1)

		if !strings.Contains(uri, domain) {
			continue
		}

		for _, keyword := range task.Keywords {
			if strings.Contains(title, keyword) {

				addSendEmailQueue(WaitEmailQueue, task.EmailTo, title, uri)

				break
			}
		}
	}
}
func hasEmailSent(queue, uri string) bool {
	r := storage.Redis()
	return r.Hexists(queue, uri)
}

func addSendEmailQueue(queue string, emailTo string, title string, uri string) {
	content := strings.Join([]string{
		emailTo,
		title,
		uri,
	}, `,`)

	r := storage.Redis()
	r.LPush(queue, content)
}

func popEmailSendQueue(queue string) (emailTo, title, uri string) {
	r := storage.Redis()
	content, err := r.RPop(queue)
	if err == nil {
		emailTo = strings.Split(content, `,`)[0]
		title = strings.Split(content, `,`)[1]
		uri = strings.Split(content, `,`)[2]
	}
	return

}

func addQueue(queue string, uri string) {
	uri = removeHashTag(uri)

	r := storage.Redis()
	r.LPush(queue, uri)
}

func popQueue(queue string) (uri string) {

	r := storage.Redis()
	uri, _ = r.RPop(queue)
	return
}

func formatUrl(uri string) string {
	uri = removeHashTag(uri)
	uri = strings.Replace(uri, "../../", "/", -1)
	return uri
}

func removeHashTag(s string) string {
	s = regex.Replace(s, `#.*`, "")
	return s
}
