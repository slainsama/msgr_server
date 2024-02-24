package handler

import (
	"sync"
	"sync/atomic"

	"github.com/smallnest/safemap"
)

var mutexPool = sync.Pool{
	New: func() interface{} {
		return &sync.Mutex{}
	},
}

type LockObj struct {
	Lock *sync.Mutex
	Num  int64
}

type KeyLock struct {
	locks *safemap.SafeMap[PersistenceKey, *LockObj]
}

func NewKeyLock() *KeyLock {
	return &KeyLock{
		locks: safemap.New[PersistenceKey, *LockObj](),
	}
}

func (l *KeyLock) getLock(key PersistenceKey) *sync.Mutex {
	if lockObj, ok := l.locks.Get(key); ok {
		atomic.AddInt64(&lockObj.Num, 1)
		return lockObj.Lock
	}
	lock := mutexPool.Get().(*sync.Mutex)
	l.locks.Set(key, &LockObj{
		Lock: lock,
	})
	return lock
}

func (l *KeyLock) Lock(key PersistenceKey) {
	l.getLock(key).Lock()
}

func (l *KeyLock) Unlock(key PersistenceKey) {
	lock, ok := l.locks.Get(key)
	if !ok {
		return
	}

	lock.Lock.Unlock()
	atomic.AddInt64(&lock.Num, -1)
	if lock.Num < 0 {
		atomic.AddInt64(&lock.Num, 1)
	}
	//clean
	for pair := range l.locks.IterBuffered() {
		if pair.Val.Num <= 0 {
			mutexPool.Put(pair.Val.Lock)
			l.locks.Remove(pair.Key)
		}
	}
}
