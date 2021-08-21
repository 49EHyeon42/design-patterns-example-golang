package singleton

import "sync"

/*
1. Bad way, thread usage is unsafe
type singleton map[string]string

var (
	instance singleton
)

func New() singleton {
	if instance == nil {
		instance = make(singleton)
	}
	return instance
}
*/

/*
2. thread usage is safe
var lock = &sync.Mutex{}

type singleton map[string]string

var (
	instance singleton
)

func New() singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = make(singleton)
	}
	return instance
}
*/

type singleton map[string]string

var (
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
