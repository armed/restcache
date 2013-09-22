package cache

import "time"

type cacheItem struct {
	duration   time.Duration
	createTime time.Time
	value      string
}

func (ci *cacheItem) isExpired() bool {
	return time.Since(ci.createTime) > ci.duration
}

func newCacheItem(value, duration string) (*cacheItem, error) {
	if duration, err := time.ParseDuration(duration); err == nil {
		return &cacheItem{duration, time.Now(), value}, nil
	} else {
		return nil, err
	}
}
