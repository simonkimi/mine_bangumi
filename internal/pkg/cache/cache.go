package cache

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)
import "github.com/patrickmn/go-cache"

var instance *cache.Cache

func init() {
	instance = cache.New(5*time.Minute, 5*time.Minute)
}

func Add(tag string, key string, value interface{}) {
	k := fmt.Sprintf("%s:%s", tag, key)
	instance.Set(k, value, cache.DefaultExpiration)
}

func Get[T any](tag string, key string) (*T, bool) {
	k := fmt.Sprintf("%s:%s", tag, key)
	if value, exist := instance.Get(k); exist {
		if v, ok := value.(T); ok {
			return &v, true
		} else {
			logrus.Warnf("Cache value type error, key: %s, expected type: %T, actual type: %T", k, new(T), value)
		}
	}
	return nil, false
}

func Delete(tag string, key string) {
	k := fmt.Sprintf("%s:%s", tag, key)
	instance.Delete(k)
}
