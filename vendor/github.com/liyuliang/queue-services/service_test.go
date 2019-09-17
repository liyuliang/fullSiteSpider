package services

import (
	"testing"
	"time"
)

func TestService(t *testing.T) {

	AddMultiProcessTask("test println hello", func(workerNum int) (err error) {
		println("hello ~")
		return err
	})
	AddMultiProcessTask("test println world", func(workerNum int) (err error) {
		println("world ~")
		return err
	})

	AddSingleProcessTask("test println --- ", func(workerNum int) (err error) {
		println("---")
		return err
	})

	// Async
	// build up 5 goroutine under a clide goroutine to run
	Service().SetWorkerNum(5).SetIsDebug(true).Start(false)
	// after 3 second, all goroutine will stop
	time.Sleep(time.Second * 3)

	// Sync
	//or you can run forever to setting isBlock:true
	//Service().SetWorkerNum(5).Start(true)
}

func TestGetServiceNames(t *testing.T) {

	multiProcessTasks = make(map[string]taskMethod)
	singleProcessTasks = make(map[string]taskMethod)

	AddMultiProcessTask("m-method 1", func(workerNum int) (err error) {
		return
	})
	AddMultiProcessTask("m-method 2", func(workerNum int) (err error) {
		return
	})
	AddMultiProcessTask("m-method 3", func(workerNum int) (err error) {
		return
	})

	AddSingleProcessTask("s-method 1", func(workerNum int) (err error) {
		return
	})

	AddSingleProcessTask("s-method 2", func(workerNum int) (err error) {
		return
	})

	names := GetMultiProcessTaskNames()
	if len(names) != 3 {
		t.Error("GetMultiProcessTaskNames should be 3,but get",len(names))
	}
	names = GetSingleProcessTaskNames()
	if len(names) != 2 {
		t.Error("GetSingleProcessTaskNames should be 2,but get",len(names))
	}
}
