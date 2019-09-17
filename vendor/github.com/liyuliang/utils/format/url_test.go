package format

import "testing"

const url = "https://www.baidu.com/"
const urlEncode = "https%3A%2F%2Fwww.baidu.com%2F"

func TestUrlEncode(t *testing.T) {

	result := UrlEncode(url)

	if result != urlEncode {
		t.Error("url encode wrong")
	}
}

func TestUrlDecode(t *testing.T) {

	result, err := UrlDecode(urlEncode)
	if err != nil {
		t.Error(err.Error())
	} else {
		if result != url {
			t.Error(result)
			t.Error("url decode wrong")
		}
	}
}
