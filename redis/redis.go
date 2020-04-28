package redis

import (
	"os"

	"github.com/go-redis/redis/v7"
)

func NewRedisClient() (*redis.Client,error) {
	url := os.Getenv("REDIS_CLIENT_URL")
	if url == "" {
		redisDB := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // use default Addr
			Password: "",               // no password set
			DB:       0,                // use default DB
		})
		
		if _, err := redisDB.Ping().Result(); err != nil {
			return nil,err
		}
		return  redisDB,nil
	}

	redisDB := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",  
		DB:       0,         
	})
	if _, err := redisDB.Ping().Result(); err != nil {
		return nil,err
	}
	return  redisDB,nil
}