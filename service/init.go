package service

import (
	"github.com/liyuliang/queue-services"
	"fullSiteSpider/storage"
	"github.com/liyuliang/utils/regex"
	"github.com/pkg/errors"
)

const WaitCrawler = "WAIT_CRAWLER"

var EmptyQueueErr = errors.New("Empty queue")

func Init() {

	indexUrl := "http://zfcj.gz.gov.cn/gzcc/index.shtml"

	addQueue(WaitCrawler, indexUrl)

	services.AddMultiProcessTask("spider run...", func(workerNum int) (err error) {

		url := popQueue(WaitCrawler)
		if url != "" {
			spiderRun(url)
			return
		} else {
			return EmptyQueueErr
		}
	})
}

func spiderRun(uri string) {

	dom, err := parseHtml(uri)
	defer func() {
		dom.Close()
	}()

	if err == nil {
		title := getTitle(dom)
		hrefs := getHrefs(dom)
	}

	r := storage.Redis()

	title, _ := r.Get(uri)

	if title == "" {

		r.Set(uri, title)

		for _, href := range hrefs {
			href = formatUrl(href)
			addQueue(WaitCrawler, href)
		}
	}
}
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
