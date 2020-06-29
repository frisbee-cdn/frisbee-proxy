package cache

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

	return nil, nil
}

// Put
func (m *MapDataStore) Put(url string, content []byte) error {

	// TODO: When add put now = time.Now() and set AddedAt
	return nil
}

// Delete
func (m *MapDataStore) Delete(url string) error {
	return nil
}

// Contains
func (m *MapDataStore) Contains(url string) bool {
	return false
}
