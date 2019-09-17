package request

import (
	"bytes"
	"errors"
	"github.com/imroc/req"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/utils/regex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetHost(s string) string {
	u, err := url.Parse(s)
	if err == nil {
		return u.Scheme + "://" + u.Host
	}
	return ""
}

func UrlRemoveHost(url string) string {
	host := GetHost(url)
	url = strings.Replace(url, host, "", -1)
	return url
}

func GetReferer(uri string) (string, error) {
	U, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	referer := regex.Get(U.Host, `([^\.]+\.[^\.]+$)`)
	referer = U.Scheme + "://www." + referer
	return referer, nil
}

func DoReq(uri string, proxy string) (resp *Response) {
	t1 := time.Now()

	resp = new(Response)

	needReferer := IsUrlNeedReferer(uri)

	header := GetHeader(uri)

	if needReferer {
		referer, err := GetReferer(uri)
		if err != nil {
			resp.Err = err
			return resp
		}
		header.Set("Referer", referer)

	}

	r := GetReq()

	if proxy != "" {
		r.SetProxyUrl("http://" + proxy)
	}

	result, err := r.Get(uri, header)
	if err != nil {
		resp.Err = err
		return resp
	}

	if result.Response().StatusCode != 200 {
		// Try again
		result, err = r.Get(uri, header)

		if err != nil {
			resp.Err = err
			return resp
		}

		if result.Response().StatusCode != 200 {
			resp.Err = errors.New("Http get request status code is not 200. ")
			return resp
		}
	}
	t2 := time.Now()

	resp.Err = err
	resp.StatusCode = result.Response().StatusCode
	resp.Data = result.String()
	resp.Speed = t2.Sub(t1).Seconds()

	result.Response().Body.Close()

	return resp
}

func IsUrlNeedReferer(uri string) bool {
	r := req.New()
	resp, err := r.Head(uri)
	if err != nil || (resp != nil && resp.Response().StatusCode == 403) {
		return true
	}
	return false
}

func GetReq() *req.Req {
	r := req.New()
	r.EnableInsecureTLS(true)
	r.SetTimeout(format.IntToTimeSecond(60))
	return r
}

func GetHeader(uri string) http.Header {
	isPc := true
	if strings.Contains(uri, "//m.") || strings.Contains(uri, "/mobile/") {
		isPc = false
	}
	header := make(http.Header)
	header.Set("User-Agent", AUserAgents(isPc))
	return header
}

func HttpGet(uri string) (httpResponse *Response) {
	resp := DoReq(uri, "")
	return resp
}

func HttpPost(uri string, params url.Values) (content string, err error) {
	resp, err := http.PostForm(uri, params)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func HttpAuthGET(account, password, uri string) (content string, err error) {
	return HttpReq("GET", account, password, uri, nil)
}

func HttpAuthPost(account, password, uri string, v url.Values) (content string, err error) {

	p := bytes.NewBufferString(v.Encode())
	return HttpReq("POST", account, password, uri, p)
}

func HttpReq(method, account, password, uri string, body io.Reader) (content string, err error) {

	req, err := http.NewRequest(method, uri, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(account, password)

	client := &http.Client{
		Timeout: time.Duration(60 * time.Second),
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return "", err
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return "", errors.New("Http response code is not 200. ")
		} else {
			bodyText, err := ioutil.ReadAll(resp.Body)
			return string(bodyText), err
		}
	}
}
