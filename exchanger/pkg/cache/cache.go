package cache

import (
	"time"
)

type Cache interface {
	Set(key string, data any, ttl time.Duration) error
	Get(key string) (value string, isExists bool)
	Delete(key string) error
	Keys(key string) []string
}
