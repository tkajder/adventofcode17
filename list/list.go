package list

import "sync"

type List struct {
	list []interface{}
	lock sync.RWMutex
}

func New() *List {
	return &List{
		list: make([]interface{}, 0),
		lock: sync.RWMutex{},
	}
}

func (l *List) Push(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list = append(l.list, value)
}

func (l *List) Pop() interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()

	if len(l.list) == 0 {
		return nil
	}

	value := l.list[0]
	l.list = l.list[1:]
	return value
}

func (l *List) Peek() interface{} {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if len(l.list) == 0 {
		return nil
	}

	return l.list[0]
}

func (l *List) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return len(l.list)
}

func (l *List) Remove(val interface{}) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	for i := 0; i < len(l.list); i++ {
		if l.list[i] == val {
			l.list = append(l.list[:i], l.list[i+1:]...)
			return true
		}
	}

	return false
}

func (l *List) Find(f func(interface{}) bool) interface{} {
	l.lock.RLock()
	defer l.lock.RUnlock()

	for i := 0; i < len(l.list); i++ {
		if f(l.list[i]) {
			return l.list[i]
		}
	}

	return nil
}

func (l *List) Init() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.list = make([]interface{}, 0)
}
