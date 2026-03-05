package config

import (
	"exchangeapp/global"
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
}
