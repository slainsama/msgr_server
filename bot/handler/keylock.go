package handler

import (
	"sync"

	"github.com/smallnest/safemap"
)

var mutexPool = sync.Pool{
	New: func() interface{} {
		return &sync.Mutex{}
	},
}

type KeyLock struct {
	locks *safemap.SafeMap[PersistenceKey, *sync.Mutex]
}

func NewKeyLock() *KeyLock {
	return &KeyLock{
		locks: safemap.New[PersistenceKey, *sync.Mutex](),
	}
}

func (l *KeyLock) Lock(key PersistenceKey) {
	var lock *sync.Mutex
	lock, ok := l.locks.Get(key)
	if !ok {
		lock = mutexPool.Get().(*sync.Mutex)
		l.locks.Set(key, lock)
	}
	lock.Lock()
}

func (l *KeyLock) Unlock(key PersistenceKey) {
	lock, ok := l.locks.Get(key)
	if ok {
		lock.Unlock()
		mutexPool.Put(lock)
		l.locks.Remove(key)
	}
}
