package pipemap

import (
	"fmt"
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

func (pm *PipeMap) Create(key string, rdr *io.PipeReader) error {
	pm.lock.Lock()
	if _, exists := pm.store[key]; exists {
		pm.lock.Unlock()
		return fmt.Errorf("key '%s' already exists", key)
	} else {
		pm.store[key] = rdr
		pm.lock.Unlock()
		return nil
	}
}

func (pm *PipeMap) Delete(key string) {
	pm.lock.Lock()
	delete(pm.store, key)
	pm.lock.Unlock()
}
