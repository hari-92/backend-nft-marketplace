package key_mutex

import "sync"

type IKeyMutex interface {
	Lock(key string)
	Unlock(key string)
}

type KeyMutex struct {
	mutexes map[string]*sync.Mutex
	mux     sync.Mutex
}

func NewKeyMutex() IKeyMutex {
	return &KeyMutex{
		mutexes: make(map[string]*sync.Mutex),
	}
}

func (k *KeyMutex) Lock(key string) {
	k.mux.Lock()
	m, ok := k.mutexes[key]
	if !ok {
		m = &sync.Mutex{}
		k.mutexes[key] = m
	}
	k.mux.Unlock()
	m.Lock()
}

func (k *KeyMutex) Unlock(key string) {
	k.mux.Lock()
	m, ok := k.mutexes[key]
	if ok {
		m.Unlock()
		delete(k.mutexes, key)
	}
	k.mux.Unlock()
}
