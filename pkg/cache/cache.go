package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type AppCache struct {
	*redis.Client
}

func NewAppCache(dsn string) (*AppCache, error) {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis dsn: %w", err)
	}
	return &AppCache{
		Client: redis.NewClient(opt),
	}, nil
}
