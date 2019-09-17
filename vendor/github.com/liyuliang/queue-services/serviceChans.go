package services

import "sync"

type serviceChan struct {
	name         string
	registerChan interface{}
}

type chanList struct {
	mutex *sync.RWMutex
	chans map[string]serviceChan
}

var list = new(chanList)

func ChanList() *chanList {

	if list.chans == nil {

		list.mutex = new(sync.RWMutex)
		list.chans = make(map[string]serviceChan)
	}

	return list
}
func (list *chanList) Names() []string {
	var names []string
	for name, _ := range list.chans {
		names = append(names, name)
	}
	return names
}

func (list *chanList) Register(name string, registerChan interface{}) {

	list.mutex.Lock()
	list.chans[name] = serviceChan{name: name, registerChan: registerChan}
	list.mutex.Unlock()
}

func (list *chanList) Get(name string) (c interface{}) {
	list.mutex.RLock()

	c = list.chans[name].registerChan
	list.mutex.RUnlock()
	return c
}
