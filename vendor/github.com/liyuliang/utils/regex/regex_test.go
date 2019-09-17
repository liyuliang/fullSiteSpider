package regex

import (
	"testing"
)

func Test_Get(t *testing.T) {
	url := "/zh-hans/slice-of-life/the-sound-of-your-heart/list?title_no=235&serviceZone=CHINA&page=11"

	title_no := Get(url, `\?title_no=(\d+)`)
	if title_no != "235" {
		t.Error("regexp get wrong , expect 235, but get", title_no)
	}
}

func TestGetAll(t *testing.T) {
	url := "js12789fnsd89wfjsf"

	nums := GetAll(url, `\d+`)
	if nums[0] != "12789" {
		t.Error("regexp get all wrong , expect 12789, but get", nums[0])
	}
	if nums[1] != "89" {
		t.Error("regexp get all wrong , expect 89, but get", nums[1])
	}
}

func TestGetAll2(t *testing.T) {
	content := `{null,"Age":24,"Balance":33.23}`
	keys := GetAll(content, `"([^\"]+)"`)
	if keys[0] != "Age" {
		t.Error("regexp get all wrong , expect Age, but get", keys[0])
	}
	if keys[1] != "Balance" {
		t.Error("regexp get all wrong , expect Balance, but get", keys[1])
	}
}

func Test_replace(t *testing.T) {
	url := "/zh-hans/slice-of-life/the-sound-of-your-heart/list?title_no=235&serviceZone=CHINA&page=11"

	url2 := Replace(url, `&page=\d+`, `&page=`+"okokok")
	url3 := Replace(url, `list\?`, "???")

	if url2 != "/zh-hans/slice-of-life/the-sound-of-your-heart/list?title_no=235&serviceZone=CHINA&page=okokok" {
		t.Error("regexp replace wrong")
	}
	if url3 != "/zh-hans/slice-of-life/the-sound-of-your-heart/???title_no=235&serviceZone=CHINA&page=11" {
		t.Error("regexp replace wrong")
	}
}
