package bad

import "sync"

type inMem struct {
	sync.Mutex
	m map[string]int
}

func NewInMem() *inMem {
	return &inMem{
		m: map[string]int{},
	}
}

func (i *inMem) Add(key string, value int) {
	i.Lock()
	defer i.Unlock()
	return i.m[key]
}

func (i *inMem) Get(key string) int {
	i.Lock()
	defer i.Unlock()
	return i.m[key]
}
