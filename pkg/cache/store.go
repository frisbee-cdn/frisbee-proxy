package cache

// Store defines general API of operations to be implemented by
// some specific store facility
type Store interface {
	Get(url string) ([]byte, error)
	Put(url string, content []byte) error
	Delete(url string) error
	Contains(url string) bool
}
