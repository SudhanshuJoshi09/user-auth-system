package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
  "sync"
  "log"
)

var (
  client *redis.Client
  ctx       context.Context
  cacheOnce sync.Once
)

func InitializeConfig(addr, password string, db, protocol int) {
  cacheOnce.Do(func() {
    client = redis.NewClient(&redis.Options{
      Addr     : addr,
      Password : password,
      DB       : db,
      Protocol : protocol,
    })
    ctx = context.Background()
  })
  return
}



func Get(key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error getting the key %s: %v", key, err)
		return "", err
	}
	return val, nil
}

func Set(key, value string) error {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error setting the key %s: %v", key, err)
		return err
	}
	return nil
}
