package cache

type Cache interface {
	Put(key, value, duration string) error
	Get(key string) (string, bool)
	Remove(key string)
}
