package service

func Init() {

	indexUrl := "http://zfcj.gz.gov.cn/gzcc/index.shtml"

	spiderRun(indexUrl)
}


func spiderRun(url string) {

	title, exist := c.Get(url)

	if !exist && title != "" {

		dom, err := parseHtml(url)
		if err == nil {
			title := getTitle(dom)
			hrefs := getHrefs(dom)

			c.Set(url, title)

			for _, href := range hrefs {
				href = formatUrl(href)
				spiderRun(href)
			}
		}
	}
}

func formatUrl(uri string) string {
	return uri
}
