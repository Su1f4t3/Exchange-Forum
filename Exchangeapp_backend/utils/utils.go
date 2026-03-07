package utils

import (
	"errors"
	"exchangeapp/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 哈希加密密码函数，接收一个字符串类型的密码作为参数，返回一个字符串类型的哈希值和一个错误对象
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

// 生成 JWT 令牌函数，接收一个字符串类型的用户名作为参数，返回一个字符串类型的 JWT 令牌和一个错误对象
func GenerateJWT(username string) (string, error) {

	// 创建一个新的 JWT 令牌对象，使用 HS256 签名方法，并设置一些声明（claims），包括用户名和过期时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// 使用 SignedString 方法对 JWT 令牌进行签名，
	signedToken, err := token.SignedString([]byte(config.AppConfig.JWT.Secret)) // secret 是用于签名的密钥

	// 返回一个字符串类型的 JWT 令牌（在签名前添加 "Bearer "（持证人） 前缀，遵循 HTTP 身份验证的标准）和一个错误对象，如果签名失败，返回一个空字符串和错误对象
	return "Bearer " + signedToken, err
}

// 比较密码哈希值函数，接收一个字符串类型的密码和一个字符串类型的哈希值作为参数，返回一个布尔值，表示密码是否与哈希值匹配
func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil // 返回一个布尔值，判断 err 是否为 nil。如果 err == nil，布尔值为 true，表示密码与哈希值匹配；如果 err != nil，布尔值为 false，表示密码与哈希值不匹配。
}

// 解析 JWT 令牌函数，接收一个字符串类型的 JWT 令牌作为参数，返回一个字符串类型的用户名和一个错误对象
func ParseJWT(tokenString string) (string, error) {

	// 去除 "Bearer " 前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// 使用 jwt.Parse 方法解析 JWT 令牌，传入一个回调函数用于验证签名方法和提供密钥，如果解析失败，返回一个空字符串和错误对象
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// 验证签名方法是否为 HMAC，如果不是，返回一个错误对象，表示不支持的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// 返回一个字节切片类型的密钥（从配置中获取 JWT 密钥），供 jwt.Parse 方法使用来验证 JWT 令牌的签名
		return []byte(config.AppConfig.JWT.Secret), nil
	})

	// 如果解析失败，返回一个空字符串和错误对象
	if err != nil {
		return "", err
	}

	// 如果解析成功，且 JWT 令牌有效，尝试从 JWT 令牌的声明（claims）中提取用户名信息，并将其存储在 username 变量中，如果提取失败，返回一个空字符串和错误对象
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)

		// 如果提取失败，返回一个空字符串和错误对象，表示 JWT 令牌中没有找到用户名信息
		if !ok {
			return "", errors.New("username not found in token")
		}

		// 返回用户名和一个 nil 错误对象，表示解析成功
		return username, nil
	}

	return "", err
}
