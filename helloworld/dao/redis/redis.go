package redis

import (
	"context"
	"strconv"

	"go.uber.org/zap"

	"github.com/go-redis/redis/v8"
	"test.com/helloworld/settings"
)

var rdb *redis.Client

func Init(redisConfig *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	err = rdb.Ping(context.Background()).Err()
	return
}

func CloseRedis() {
	if err := rdb.Close(); err != nil {
		zap.L().DPanic("fail to close redis connection", zap.Error(err))
	}
}
