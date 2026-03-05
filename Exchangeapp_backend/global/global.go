package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// 定义全局变量 Db 和 RedisDB，分别用于存储数据库连接对象和 Redis 客户端对象，以便在整个应用程序中共享和使用这些资源
var (
	Db      *gorm.DB
	RedisDB *redis.Client
)
