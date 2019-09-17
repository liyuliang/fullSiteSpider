package services

import (
	"testing"
)

func Test_Chans(t *testing.T) {

	type book struct {
		name string
	}

	ChanList().Register("testString", make(chan string))
	ChanList().Register("testString2", make(chan string, 2))
	ChanList().Register("testInt", make(chan int))
	ChanList().Register("testInt2", make(chan int, 2))

	ChanList().Register("testStruct", make(chan book))
	ChanList().Register("testStruct2", make(chan book, 2))

	c := ChanList().Get("testInt").(chan int)
	go func() {
		c <- 1
	}()
	if 1 != <-c {
		t.Error("get wrong value from a int chan")
	}

	c = ChanList().Get("testInt2").(chan int)
	go func() {
		c <- 1
	}()
	if 1 != <-c {
		t.Error("get wrong value from a int chan 2")
	}

	c2 := ChanList().Get("testStruct").(chan book)
	go func() {

		c2 <- book{name: "golang"}
	}()
	s := <-c2
	if "golang" != s.name {
		t.Error("get wrong value from a int chan 2")
	}
}
