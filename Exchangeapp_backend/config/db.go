package config

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接函数 InitDB
func InitDB() {
	// 从全局变量 AppConfig 中获取数据库连接字符串 DSN
	dsn := AppConfig.Database.Dsn

	// 使用 GORM 的 mysql 驱动打开数据库连接，如果出错则记录日志并退出程序
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to initialize database, got error: %v", err)
	}

	// 获取底层的 sql.DB 对象，并设置连接池参数，如果出错则记录日志并退出程序
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("failed to get sql.DB, got error: %v", err)
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns) // 最大空闲连接数
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns) // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)                    // 连接最大生命周期

	// 将数据库连接对象 db 赋值给全局变量 global.Db，供其他包使用
	global.Db = db

	// 使用 AutoMigrate 方法进行自动迁移，确保数据库中存在与 User、Article 和 ExchangeRate 结构体对应的表，如果迁移失败，记录日志并退出程序
	err = global.Db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.ExchangeRate{},
	) // 执行后，数据库会生成 users、articles 和 exchange_rates 三张表，分别用于存储用户信息、文章信息和汇率信息

	if err != nil {
		log.Fatalf("failed to migrate database, got error: %v", err)
	}
}
