package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	instance *cache.Cache
}

func NewCache() *Cache {
	return &Cache{
		instance: cache.New(5*time.Minute, 5*time.Minute),
	}
}

func (c *Cache) Add(tag string, key string, value interface{}) {
	k := fmt.Sprintf("%s:%s", tag, key)
	c.instance.Set(k, value, cache.DefaultExpiration)
}

func (c *Cache) Get(tag string, key string) (any, bool) {
	k := fmt.Sprintf("%s:%s", tag, key)
	if value, exist := c.instance.Get(k); exist {
		return value, true
	}
	return nil, false
}

func (c *Cache) Delete(tag string, key string) {
	k := fmt.Sprintf("%s:%s", tag, key)
	c.instance.Delete(k)
}
