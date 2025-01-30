package persistence

import "github.com/t-shah02/mochi/internal/models"

type MemoryManager struct {
	simpleKey *models.SimpleKeyValueStore
	kLists    *models.KeyToListStore
	kSets     *models.KeyToSetStore
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		simpleKey: models.NewSimpleKeyValueStore(),
		kLists:    models.NewKeyToList(),
		kSets:     models.NewKeyToSetStore(),
	}
}
