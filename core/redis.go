package core

import (
	"github.com/go-redis/redis"
	"log"
	"strings"
	"time"
)

var Redis *redis.Client

func initRedis() {
	parts := strings.SplitN(Config.Redis, "//", 2)
	Redis = redis.NewClient(&redis.Options{
		Addr:     parts[0],
		Password: parts[1],
		DB:       0, // use default DB
	})
}

func GetFromRedis(key string) string {
	v, e := Redis.Get(key).Result()
	if e != nil && e != redis.Nil {
		log.Println(e.Error())
	}
	return v
}

func SetToRedis(key, value string) {
	Redis.Set(key, value, 24*time.Hour)
}
