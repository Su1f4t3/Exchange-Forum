package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册函数，接收一个 Gin 上下文对象作为参数，处理用户注册的 HTTP 请求
func Register(ctx *gin.Context) {

	// 定义一个 User 结构体变量 user，用于接收请求中的 JSON 数据并进行绑定
	var user models.User

	// 使用 ShouldBindJSON 方法将请求中的 JSON 数据绑定到 user 变量上，如果绑定失败，返回一个 HTTP 400 错误响应，包含错误信息
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用 HashPassword 方法对 user 变量中的 Password 字段进行哈希加密，并将结果存储在 hashedPwd 变量中，如果加密失败，返回一个 HTTP 500 错误响应，包含错误信息
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将 user 变量中的 Password 字段替换为加密后的哈希值 hashedPwd，以确保在数据库中存储的密码是安全的，防止明文密码泄露
	user.Password = hashedPwd

	// 使用 GerenateJWT 方法生成一个 JWT 令牌，传入 user 变量中的 Username 字段作为参数，并将结果存储在 token 变量中，如果生成失败，返回一个 HTTP 500 错误响应，包含错误信息
	token, err := utils.GerenateJWT(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// // 使用 AutoMigrate 方法进行自动迁移，确保数据库中存在与 User 结构体对应的表，如果迁移失败，返回一个 HTTP 500 错误响应，包含错误信息
	// if err := global.Db.AutoMigrate(&user); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 使用 Create 方法将 user 变量中的数据插入数据库的 users 表中，如果插入失败，返回一个 HTTP 500 错误响应，包含错误信息
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回一个 HTTP 200 成功响应，包含生成的 JWT 令牌
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// 登录函数，接收一个 Gin 上下文对象作为参数，处理用户登录的 HTTP 请求
func Login(ctx *gin.Context) {

	// 定义一个结构体变量 input，包含两个字段 Username 和 Password，分别用于接收请求中的用户名和密码数据，并且都设置为必填字段
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 使用 ShouldBindJSON 方法将请求中的 JSON 数据绑定到 input 变量上，如果绑定失败，返回一个 HTTP 400 错误响应，包含错误信息
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 定义一个 User 结构体变量 user，用于存储从数据库中查询到的用户数据
	var user models.User

	// 使用 Where 方法从数据库中的 users 表查询第一条记录，条件是 username 字段等于 input 变量中的 Username 字段，并将结果存储在 user 变量中，如果查询失败，返回一个 HTTP 401 错误响应，包含错误信息 "Invalid username or password"
	// 这里使用了参数化查询，防止 SQL 注入攻击
	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 使用 CheckPasswordHash 方法比较 input 变量中的 Password 字段和 user 变量中的 Password 字段（哈希值），如果比较失败，返回一个 HTTP 401 错误响应，包含错误信息 "Invalid username or password"
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 使用 GerenateJWT 方法生成一个 JWT 令牌，传入 user 变量中的 Username 字段作为参数，并将结果存储在 token 变量中，如果生成失败，返回一个 HTTP 500 错误响应，包含错误信息
	token, err := utils.GerenateJWT(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回一个 HTTP 200 成功响应，包含生成的 JWT 令牌
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
