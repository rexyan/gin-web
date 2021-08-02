package redis

import (
	"fmt"
	"web_app/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(redisConfig *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisConfig.Host,
			redisConfig.Port,
		),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
