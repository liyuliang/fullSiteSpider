package format

import "testing"

func TestSliceDelDuplicate(t *testing.T) {

	input := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}
	expect := []string{"hello", "world", "golang", "ruby", "php", "java"}

	result := SliceDelDuplicate(input)

	for i, value := range result {
		if value != expect[i] {
			t.Error("SliceDelDuplicate failed, expect :" + expect[i] + " ,but get " + value)
		}
	}
}
