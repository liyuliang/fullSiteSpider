package service

import (
	"github.com/liyuliang/queue-services"
	"fullSiteSpider/storage"
	"github.com/liyuliang/utils/regex"
	"github.com/pkg/errors"
	"fullSiteSpider/config"
	"log"
	"time"
	"net/url"
	"strings"
)

const WaitCrawlerQueue = "WAIT_CRAWLER_QUEUE"
const WaitEmailQueue = "WAIT_SEND_EMAIL_QUEUE"
const EmailSendedList = "EMAIL_SEND_LIST"

var EmptyQueueErr = errors.New("Empty queue")
var ThirdPartyErr = errors.New("third party url")

var taskDomains []string

func Init() {

	initTasks()
	initTaskDomains()

	services.AddMultiProcessTask("spider run...", func(workerNum int) (err error) {

		url := popQueue(WaitCrawlerQueue)
		if url != "" {
			spiderRun(url)
			return

		} else {
			log.Println(EmptyQueueErr.Error())
			time.Sleep(time.Hour * 10)
			return EmptyQueueErr
		}
	})

	services.AddSingleProcessTask("send email...", func(workerNum int) (err error) {

		mailTo, title, uri := popEmailSendQueue()
		if mailTo != "" && title != "" && uri != "" {
			content := toMailContent(title, uri)

			err = sendEmail(mailTo, content)

			if err == nil {
				recordEmailSended(uri, title)
			}
		}

		return err
	})
}
func initTaskDomains() {

	for _, task := range config.Tasks() {
		u, _ := url.Parse(task.Url)
		taskDomains = append(taskDomains, u.Hostname())
	}
}

func initTasks() {
	for _, task := range config.Tasks() {
		addQueue(WaitCrawlerQueue, task.Url)
	}
}

func spiderRun(uri string) error {

	if isThirdPartyUrl(uri) {
		return ThirdPartyErr
	}

	dom, err := parseHtml(uri)
	defer func() {
		dom.Close()
	}()

	if err != nil {
		return err
	}

	title := getTitle(dom)
	hrefs := getHrefs(dom)

	r := storage.Redis()

	titleRecord, _ := r.Get(uri)

	if titleRecord == "" {

		checkAndSendEmail(uri, title)

		for _, href := range hrefs {
			href = formatUrl(href)
			addQueue(WaitCrawlerQueue, href)
		}
	}
}

func isThirdPartyUrl(uri string) bool {

	isThirdParty := true
	for _, domain := range taskDomains {
		if strings.Contains(uri, domain) {
			isThirdParty = false
			break
		}
	}
	return isThirdParty
}

func checkAndSendEmail(uri string, title string) {
	if hasEmailSended(uri) {
		return
	}

	for name, task := range config.Tasks() {
		domain := strings.Replace(name, "task.", "", -1)

		if !strings.Contains(uri, domain) {
			continue
		}

		for _, keyword := range task.Keywords {
			if strings.Contains(title, keyword) {

				addSendEmailQueue(task.EmailTo, title, uri)

				break
			}
		}
	}
}

func addSendEmailQueue(emailTo string, title string, uri string) {

}

func popEmailSendQueue() (emailTo, title, uri string) {

}

func addQueue(queue string, uri string) {
	uri = removeHashTag(uri)

}

func popQueue(queue string) (uri string) {

}

func formatUrl(uri string) string {
	uri = removeHashTag(uri)
	return uri
}

func removeHashTag(s string) string {
	s = regex.Replace(s, `#.*`, "")
	return s
}
