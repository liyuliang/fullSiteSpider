package format

import (
	URL "net/url"
	"errors"
	"strings"
)

func UrlDecode(url string) (string, error) {
	if url == "" {
		return "", errors.New("Url decode miss argument url")
	}

	url, err := URL.QueryUnescape(url)
	return url, err
}

func UrlEncode(url string) string {
	if url != "" {
		url = URL.QueryEscape(url)
	}
	return url
}

func UrlFormat(url string) (string) {
	if !strings.Contains(url, "http:") || !strings.Contains(url, "https:") {
		return "http://" + url
	} else {
		return url
	}
}
