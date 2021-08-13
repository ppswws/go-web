package initialize

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var (
	redisIns  *redis.Client
	onceRedis sync.Once
)

// 初始化Redis
func Redis() {

	// 单例模式，保证生命周期只初始化一次
	onceRedis.Do(func() {
		redisIns = redis.NewClient(&redis.Options{
			Addr:     configuration.Redis.Addr,
			Password: configuration.Redis.Password,
			DB:       configuration.Redis.Db,
		})
	})
	err := redisIns.Ping().Err()
	if err != nil || redisIns == nil {
		panic(fmt.Sprintf("初始化Redis异常：%v", err))
	}

	fmt.Println("初始化Redis完成...")
}

// 暴露给其他包
func GetRedisIns() *redis.Client {
	return redisIns
}