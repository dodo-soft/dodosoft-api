package events

import (
	"time"
)

type cache struct {
	entries map[string]*entry
}

type entry struct {
	timeToLive int64
	value      interface{}
}

func newCache() *cache {
	c := cache{}
	c.entries = map[string]*entry{}
	return &c
}

func (c *cache) get(key string) *interface{} {
	e := c.entries[key]
	if e == nil {
		return nil
	}
	if time.Now().UnixNano() > e.timeToLive {
		delete(c.entries, key)
		return nil
	}
	return &e.value
}

func (c *cache) Put(key string, value interface{}) {
	ttl := time.Now().UnixNano() + time.Hour.Nanoseconds()
	e := entry{ttl, value}
	c.entries[key] = &e
}
