package cache

import (
	"errors"
	"time"
)

// ErrNotFound marks a not found error
var ErrNotFound = errors.New("Not Found")

// MapDataStore implements a simple key-value pair store with web resource
// kept in memory
type MapDataStore struct {
	Store

	resources map[string]WebResource
	capacity  uint
}

// NewMapDataStore creates a new MapDataStore
func NewMapDataStore(c uint) *MapDataStore {

	return &MapDataStore{
		resources: make(map[string]WebResource),
		capacity:  c,
	}
}

// Get
func (m *MapDataStore) Get(url string) ([]byte, error) {

	res, found := m.resources[url]
	if !found {
		return nil, ErrNotFound
	}
	return res.Content, nil
}

// Put
func (m *MapDataStore) Put(url string, content []byte) error {
	nwr := WebResource{
		Content: content,
		AddedAt: time.Now(),
	}
	m.resources[url] = nwr
	return nil

	// TODO: check with capacity, if excedes capacity, then remove the least requested ones
}

// Delete
func (m *MapDataStore) Delete(url string) error {
	delete(m.resources, url)
	return nil
}

// Contains
func (m *MapDataStore) Contains(url string) bool {
	_, found := m.resources[url]
	return found
}
