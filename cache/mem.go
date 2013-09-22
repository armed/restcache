package cache

import (
	"sync"
	"time"
)

type memCache struct {
	sync.Mutex
	defaultDuration string
	items           map[string]*cacheItem
}

func (c *memCache) Put(key, value, duration string) error {
	if duration == "" {
		duration = c.defaultDuration
	}

	item, err := newCacheItem(value, duration)

	if err != nil {
		return err
	} else {
		c.Lock()
		c.items[key] = item
		c.Unlock()
		return nil
	}
}

func (c *memCache) Get(key string) (string, bool) {
	if item, ok := c.items[key]; ok {
		return item.value, true
	} else {
		return "", false
	}
}

func (c *memCache) Remove(key string) {
	delete(c.items, key)
}

func (c *memCache) cleanExpired() {
	c.Lock()
	for k, v := range c.items {
		if v.isExpired() {
			c.Remove(k)
		}
	}
	c.Unlock()
}

func New(defaultDuration string) Cache {
	c := &memCache{
		defaultDuration: defaultDuration,
		items:           map[string]*cacheItem{},
	}

	go func() {
		tick := time.Tick(10 * time.Second)
		for _ = range tick {
			c.cleanExpired()
		}
	}()

	return c
}
