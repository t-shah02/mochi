package models

import (
	"fmt"
	"strings"
	"sync"
)

type KeyToSetStore struct {
	data sync.Map
}

func NewKeyToSetStore() *KeyToSetStore {
	return &KeyToSetStore{
		data: sync.Map{},
	}
}

func (store *KeyToSetStore) Set(key string, value string) bool {
	processedValues := strings.Split(strings.Trim(value, "{}"), ",")

	store.data.Delete(key)

	newSet := &sync.Map{}
	store.data.Store(key, newSet)

	for processedValue := range processedValues {
		newSet.Store(processedValue, true)
	}

	return true
}

func (store *KeyToSetStore) Del(key string) bool {
	store.data.Delete(key)
	return true
}

func (store *KeyToSetStore) AddItemToKey(key string, newValue string) bool {
	keyMembers, found := store.data.Load(key)
	if !found {
		newSet := sync.Map{}
		newSet.Store(newValue, true)
		store.data.Store(key, &newSet)

		return true
	}

	keyMembers.(*sync.Map).Store(newValue, true)
	return true
}

func (store *KeyToSetStore) RemoveItemFromKey(key string, targetValue string) bool {
	keyMembers, found := store.data.Load(key)
	if !found {
		return false
	}

	keyMembers.(*sync.Map).Delete(targetValue)
	return true
}

func (store *KeyToSetStore) GetMembers(key string) (string, bool) {
	keyMembers, found := store.data.Load(key)
	if !found {
		return "", false
	}

	var result []string
	keyMembers.(*sync.Map).Range(func(key, _ any) bool {
		result = append(result, key.(string))
		return true
	})

	stringifiedResult := fmt.Sprintf("{%s}", strings.Join(result, ","))
	return stringifiedResult, true
}
