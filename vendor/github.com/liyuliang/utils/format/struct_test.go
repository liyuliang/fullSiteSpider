package format

import "testing"

func TestClear(t *testing.T) {

	type People struct {
		Name    string
		age     int
		address string
		Money   float32
	}

	me := new(People)
	me.Name = "liang"
	me.address = "China guangdong"
	me.age = 25
	me.Money = 102.5

	ClearStruct(me)
	if me.Name!= "" && me.address != "" && me.age != 0 && me.Money != 0{
		t.Error("the method ClearStruct has something wrong")
	}
}
