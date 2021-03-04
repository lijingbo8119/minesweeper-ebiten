package cursor

import "sync"

type valEvent func(val interface{})

type Val struct {
	val    interface{}
	mut    sync.RWMutex
	events []valEvent
}

func (this *Val) Get() interface{} {
	this.mut.RLock()
	defer this.mut.RUnlock()
	return this.val
}

func (this *Val) Set(val interface{}) *Val {
	this.mut.Lock()
	defer this.mut.Unlock()
	this.val = val
	for _, e := range this.events {
		e(this.val)
	}
	return this
}

func (this *Val) Event(e valEvent) *Val {
	this.events = append(this.events, e)
	return this
}

func NewVal(val interface{}) *Val {
	return &Val{val: val}
}
