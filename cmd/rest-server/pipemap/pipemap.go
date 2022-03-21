package pipemap

import (
	"io"
	"sync"
)

type PipeMap struct {
	store map[string]*io.PipeReader
	lock  sync.Mutex
}

func New() *PipeMap {
	return &PipeMap{
		store: make(map[string]*io.PipeReader),
		lock:  sync.Mutex{},
	}
}

func (pm *PipeMap) Get(key string) (*io.PipeReader, bool) {
	pm.lock.Lock()
	rdr, ok := pm.store[key]
	pm.lock.Unlock()

	return rdr, ok
}

func (pm *PipeMap) Create(key string, rdr *io.PipeReader) {
	pm.lock.Lock()
	pm.store[key] = rdr
	pm.lock.Unlock()
}

func (pm *PipeMap) Delete(key string) {
	pm.lock.Lock()
	delete(pm.store, key)
	pm.lock.Unlock()
}
