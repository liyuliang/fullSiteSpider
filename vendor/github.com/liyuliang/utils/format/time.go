package format

import "time"

func IntToTime(content int) time.Duration {
	return time.Duration(content)
}

func TimeToInt(t time.Time) int {
	return int(t.Unix())
}

func IntToTimeSecond(content int) time.Duration {
	return IntToTime(content) * time.Second
}

func SecondTimeAfterNow(second int) time.Time {
	if second == 0 {
		return time.Now()
	}else{
		return time.Now().Add(+ IntToTimeSecond(second))
	}
}
