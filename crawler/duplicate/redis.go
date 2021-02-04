package duplicate

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

const redis_key = "fetched_url"

var ctx = context.Background()

type RedisDuplicator struct {
	Client *redis.Client
}

func NewRedisDuplicator(host string, password string) *RedisDuplicator {
	// "127.0.0.1:6379"

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	return &RedisDuplicator{Client: client}
}

func (d *RedisDuplicator) Duplicate(url string) bool {
	ok, err := d.Client.HExists(ctx, redis_key, url).Result()
	if err != nil {
		log.Printf("Duplicate: error url %s: %v", url, err)
		return false
	}
	if ok {
		log.Printf("Duplicate: duplicated url %s", url)
		return true
	} else { // 保存到Redis
		d.Client.HSet(ctx, redis_key, url, "ok")
		return false
	}
}
