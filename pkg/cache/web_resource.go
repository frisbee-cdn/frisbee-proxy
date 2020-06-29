package cache

import "time"

// WebResource defines the web resource to be stored in the cache
type WebResource struct {
	Content []byte
	AddedAt time.Time
}

// TODO: Implement Sort, Len, Swap, Less to sort web resources
