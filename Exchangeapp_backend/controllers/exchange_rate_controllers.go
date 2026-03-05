package controllers

import (
	"errors"
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建汇率函数，接收一个 Gin 上下文对象作为参数，处理创建汇率的 HTTP 请求，首先将请求体中的 JSON 数据绑定到一个 ExchangeRate 结构体实例中，如果绑定失败，返回一个 HTTP 400 错误响应，包含错误信息
func CreateExchangeRate(ctx *gin.Context) {

	// 定义一个 ExchangeRate 结构体变量 exchangeRate，用于存储从请求体中绑定的汇率数据
	var exchangeRate models.ExchangeRate

	// 使用 ShouldBindJSON 方法将请求体中的 JSON 数据绑定到 exchangeRate 变量上，如果绑定失败，返回一个 HTTP 400 错误响应，包含错误信息
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 将当前的时间赋值给 exchangeRate 结构体的 Time 字段，以记录汇率数据的创建时间
	exchangeRate.Time = time.Now()

	// 使用 AutoMigrate 方法进行自动迁移，确保数据库中存在与 ExchangeRate 结构体对应的表，如果迁移失败，返回一个 HTTP 500 错误响应，包含错误信息
	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 使用 Creat 方法将 exchangeRate 变量中的数据插入数据库的 exchange_rates 表中，如果插入失败，返回一个 HTTP 500 错误响应，包含错误信息
	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回一个 HTTP 200 成功响应，包含创建成功的汇率数据
	ctx.JSON(http.StatusOK, exchangeRate)
}

// 获取汇率列表函数，接受一个 Gin 上下文对象作为参数，处理获取汇率列表的 HTTP 请求，首先定义一个 ExchangeRate 结构体切片变量 exchangeRates，用于存储从数据库中查询到的汇率列表，然后使用 Find 方法从数据库中的 exchange_rates 表查询所有的汇率数据，并将结果存储在 exchangeRates 变量中，如果查询失败，且错误类型是 gorm.ErrRecordNotFound，表示没有找到任何汇率数据，此时返回一个 HTTP 404 错误响应，包含错误信息 "No exchange rates found"，如果查询失败但错误类型不是 gorm.ErrRecordNotFound，表示查询过程中发生了其他错误，此时返回一个 HTTP 500 错误响应，包含错误信息 "Failed to retrieve exchange rates"，如果查询成功，则返回一个 HTTP 200 成功响应，包含查询到的汇率列表数据
func GetExchangeRates(ctx *gin.Context) {

	// 定义一个 ExchangeRate 结构体切片变量 exchangeRates，用于存储从数据库中查询到的汇率列表
	var exchangeRates []models.ExchangeRate

	// 使用 Find 方法从数据库中的 exchange_rates 表查询所有的汇率数据，并将结果存储在 exchangeRates 变量中，如果查询失败，且错误类型是 gorm.ErrRecordNotFound，表示没有找到任何汇率数据，此时返回一个 HTTP 404 错误响应，包含错误信息 "No exchange rates found"，如果查询失败但错误类型不是 gorm.ErrRecordNotFound，表示查询过程中发生了其他错误，此时返回一个 HTTP 500 错误响应，包含错误信息 "Failed to retrieve exchange rates"，如果查询成功，则返回一个 HTTP 200 成功响应，包含查询到的汇率列表数据
	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No exchange rates found"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve exchange rates"})
			return
		}
	}

	// 返回一个 HTTP 200 成功响应，包含查询到的汇率列表数据
	ctx.JSON(http.StatusOK, exchangeRates)
}
