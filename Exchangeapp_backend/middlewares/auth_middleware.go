package middlewares

import (
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 是一个 Gin 中间件函数，用于验证请求中的 JWT 令牌，确保只有经过身份验证的用户才能访问受保护的 API 路径
func AuthMiddleware() gin.HandlerFunc {
	// 返回一个匿名函数，接收一个 Gin 上下文对象作为参数，处理每个 HTTP 请求的身份验证逻辑
	return func(ctx *gin.Context) {

		// 从请求头中获取 Authorization 字段的值，存储在 token 变量中，如果 token 为空，返回一个 HTTP 401 错误响应，包含错误信息 "Authorization
		// 1. 检查 token 是否存在
		token := ctx.GetHeader("Authorization")

		// 如果 token 为空，返回一个 HTTP 401 错误响应，包含错误信息 "Authorization token is required"，并且调用 ctx.Abort() 方法终止请求的处理，防止继续执行后续的处理函数
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			ctx.Abort()
			return
		}

		// 如果 token 不为空，调用 utils.ParseJWT 方法解析 token，获取其中的用户名信息，并将结果存储在 username 变量中，如果解析失败，返回一个 HTTP 401 错误响应，包含错误信息 "Invalid token: " 加上具体的错误信息，并且调用 ctx.Abort() 方法终止请求的处理
		// 2. 解析 token，获取用户名信息，失败就 Abort() ，禁止入门，终止请求的处理
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			ctx.Abort() // 终止请求的处理，防止继续执行后续的处理函数
			return
		}

		ctx.Set("username", username) // 将解析出的用户名信息存储在 Gin 上下文中，以便后续的处理函数可以使用 ctx.Get("username") 来获取当前请求的用户名信息
		ctx.Next()                    // 调用 ctx.Next() 方法继续执行后续的处理函数，允许请求继续向下传递，允许 controller 访问到 username 信息
	}
}
