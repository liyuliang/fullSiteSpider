package cli

import (
	"os"
)

type param string

func (p param) ToString() string {
	return string(p)
}

func (p param) IsExist() bool {
	return p.ToString() != ""
}

func GetParam(paramIndex int) (arg param) {

	argsNum := len(os.Args)
	if argsNum > paramIndex {
		arg = param(os.Args[paramIndex])
	}
	return arg
}

func GetAllParams() (args []param) {

	argsNum := len(os.Args)
	if argsNum > 1 {

		for _, value := range os.Args[1:argsNum] {
			args = append(args, param(value))
		}
	}

	return args
}
