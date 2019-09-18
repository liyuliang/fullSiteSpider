package storage

import (
	"testing"
	"github.com/liyuliang/utils/format"
)

const REDIS_TEST_KEY string = "REDIS_TEST_KEY"

func TestRedisCache_Del_Start(t *testing.T) {

	Redis().do("del", REDIS_TEST_KEY)
}

func TestRedisCache_ZAdd(t *testing.T) {
	for i := 1; i < 5; i++ {
		Redis().ZAdd(REDIS_TEST_KEY, i, format.IntToStr(i)+"_val")
	}
}

func TestRedisCache_ZRange(t *testing.T) {

	values, err := Redis().ZRangeByScore(REDIS_TEST_KEY, "3", "4")
	if err != nil {
		t.Error(err.Error())
	}

	if len(values) != 2 {
		t.Error("ZRangeByScore should get 2 value, but get ", len(values))
	}
	if "3_val" != values[0] {
		t.Error("ZRangeByScore value wrong ,should be 3_val, but ", values[0])
	}
	if "4_val" != values[1] {
		t.Error("ZRangeByScore value wrong ,should be 4_val, but ", values[1])
	}

	values, err = Redis().ZRangeByScore(REDIS_TEST_KEY, "(3", "4")
	if err != nil {
		t.Error(err.Error())
	}

	if len(values) != 1 {
		t.Error("ZRangeByScore should get 1 value, but get ", len(values))
	}
	if "4_val" != values[0] {
		t.Error("ZRangeByScore value wrong ,should be 3_val, but ", values[0])
	}
}

func TestRedisCache_ZRangeByScore_Null(t *testing.T) {
	values, _ := Redis().ZRangeByScore(REDIS_TEST_KEY, "30", "40")

	for _, value := range values {
		if value != "" {
			t.Error("zrangebyscore over max index should get null ,but ", value)
		}
	}

}

func TestRedisCache_Del_End(t *testing.T) {

	Redis().do("del", REDIS_TEST_KEY)
}
