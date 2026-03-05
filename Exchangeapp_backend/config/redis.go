package config

import (
	"exchangeapp/global"

	"github.com/go-redis/redis"
)

// 定义一个全局变量 RedisClient，类型是指向 redis.Client 结构体的指针
var RedisClient *redis.Client

// 初始化 Redis 连接函数 InitRedis
func InitRedis() {
	// 使用 go-redis 包创建一个新的 Redis 客户端，连接参数从全局变量 AppConfig 中获取
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr, // Redis 服务器地址
		DB:       0,                    // 使用默认的 Redis 数据库
		Password: "",                   // 没有密码
	})

	// RedisClient.Ping() 发送一个 PING 命令到 Redis 服务器，返回 *StatusCmd 对象，调用 Result() 方法获取结果，如果出错则 panic 报错
	_, err := RedisClient.Ping().Result() // .Result() 只返回结果和错误，忽略其他信息
	if err != nil {
		panic("Failed to connect to Redis")
	}

	// 将 Redis 客户端对象 RedisClient 赋值给全局变量 global.RedisDB，供其他包使用
	global.RedisDB = RedisClient
}
