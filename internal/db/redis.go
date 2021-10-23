package db

import "github.com/go-redis/redis"

func New(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: cfg.Address, Password: cfg.Password})
}
