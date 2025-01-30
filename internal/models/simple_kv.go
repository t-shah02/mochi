package models

import "sync"

type SimpleKeyValueStore struct {
	data sync.Map
}

func NewSimpleKeyValueStore() *SimpleKeyValueStore {
	return &SimpleKeyValueStore{
		data: sync.Map{},
	}
}

func (store *SimpleKeyValueStore) Get(key string) (string, bool) {
	value, found := store.data.Load(key)
	if !found {
		return "", false
	}

	return value.(string), true
}

func (store *SimpleKeyValueStore) Set(key string, value string) bool {
	store.data.Store(key, value)

	return true
}

func (store *SimpleKeyValueStore) Del(key string) bool {
	store.data.Delete(key)
	return true
}
