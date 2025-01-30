package models

type KeyToListStore struct {
	data map[string][]string
}

func NewKeyToList() *KeyToListStore {
	return &KeyToListStore{
		data: make(map[string][]string),
	}
}
