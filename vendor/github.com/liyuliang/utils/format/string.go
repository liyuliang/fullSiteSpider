package format

import (
	"strconv"
	"math/big"
	"net"
	"errors"
	"fmt"
)

func IntToStr(content int) string {
	return strconv.Itoa(content)
}

func Int64ToStr(content int64) string {
	return strconv.FormatInt(content, 10)
}

func StrToInt(content string) int {
	result, err := strconv.Atoi(content)
	if nil != err {
		return 0
	} else {
		return result
	}
}

func StrToInt64(content string) int64 {
	i, err := strconv.ParseInt(content, 10, 64)
	if err != nil {
		return 0
	} else {
		return i
	}
}

func IpToInt64(ipAddr string) (int64, error) {
	ret := big.NewInt(0)
	i := net.ParseIP(ipAddr).To4()
	if i != nil {
		ret.SetBytes(i)
		return ret.Int64(), nil
	}
	return 0, errors.New("Ip address not available. ")
}

func Int64ToIp(ipInt int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ipInt>>24), byte(ipInt>>16), byte(ipInt>>8), byte(ipInt))
}
