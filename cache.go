package main

import (
	"sync"
)

// 该文件为lru添加并发特性
type cache struct {
	mu         sync.Mutex
	lru        *Cache
	cacheBytes int64
}

func (c *cache) update(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = New(c.cacheBytes, nil)
	}

	c.lru.Update(key, value)

}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}