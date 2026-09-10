package engine

import (
	"sync"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("key not found")

type Storage interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Delete(key string) error
}

type StorageEngine struct {
	mu   sync.RWMutex
	data map[string]any
}

func NewStorageEngine() Storage {
	return &StorageEngine{
		mu:   sync.RWMutex{},
		data: make(map[string]any),
	}
}

func (e *StorageEngine) Set(key string, value any) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.data[key] = value
	return nil
}

func (e *StorageEngine) Get(key string) (any, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	v, ok := e.data[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (e *StorageEngine) Delete(key string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	_, ok := e.data[key]
	if !ok {
		return ErrNotFound
	}
	delete(e.data, key)
	return nil
}
