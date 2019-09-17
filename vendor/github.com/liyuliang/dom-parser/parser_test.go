package parser

import (
	"testing"
)

func Test_Find_One_Dom(t *testing.T) {

	html := "<body><div>12344<a href='11111'>11111q</a><a href='22222'>22222q</a></div></body>"
	dom, err := InitDom(html)

	if err != nil {
		t.Error(err.Error())
	}else {
		text := dom.Find("div").Text()
		if "1234411111q22222q" != text {
			t.Error("goquery find div node ,get all text failed")
		}

		href, is_find := dom.Find("a").Attr("href")
		if !is_find {
			t.Error("goquery find one a href failed")
		}

		if "11111" != href {
			t.Error("goquery get one a href value failed")
		}
	}

}


func Test_Find_All_Dom(t *testing.T) {

	html := "<body><div>12344<a href='11111'>11111q</a><a href='22222'>22222q</a></div></body>"
	dom, err := InitDom(html)

	if err != nil {
		t.Error(err.Error())
	}else {
		doms := dom.FindAll("a")
		if 2 != len(doms) {
			t.Error("goquery find all a count wrong")
		}
		if value1, _ := doms[0].Attr("href"); value1 != "11111" {
			t.Error("goquery find all a , first element wrong")
		}
		if value2, _ := doms[1].Attr("href"); value2 != "22222" {
			t.Error("goquery find all a , second element wrong")
		}
	}
}

func Test_LastOne_Dom(t *testing.T) {
	html := "<body><div>12344<a href='11111'>11111q</a><a href='22222'>22222q</a></div></body>"
	dom, err := InitDom(html)

	if err != nil {
		t.Error(err.Error())
	}else {
		doms := dom.FindAll("a")

		lastA := doms[(len(doms) - 1)]

		lastone_ele, _ := lastA.Attr("href")
		if "22222" != lastone_ele {
			t.Error("goquerey get lastone element value failed")
		}
	}
}
