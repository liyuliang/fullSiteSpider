package services

type taskMethod func(workerNum int) (err error)

var multiProcessTasks map[string]taskMethod
var singleProcessTasks map[string]taskMethod

func init() {
	multiProcessTasks = make(map[string]taskMethod)
	singleProcessTasks = make(map[string]taskMethod)
}

func AddMultiProcessTask(methodName string, method taskMethod) {

	multiProcessTasks[methodName] = method
}

func AddSingleProcessTask(methodName string, method taskMethod) {
	singleProcessTasks[methodName] = method
}

func GetMultiProcessTaskNames() (names []string) {
	for name, _ := range multiProcessTasks {
		names = append(names, name)
	}
	return names
}

func GetSingleProcessTaskNames() (names []string) {
	for name, _ := range singleProcessTasks {
		names = append(names, name)
	}
	return names
}
