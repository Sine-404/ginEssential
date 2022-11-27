package middel

import (
	"WebProject/ginEssential/mapping"
	"WebProject/ginEssential/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 定义gin的中间键
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization(授权) header
		tokenString := ctx.GetHeader("Authorization")

		//进行校验
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort() //阻止调用后续函数
			return
		}

		//截取token有用位置
		tokenString = tokenString[7:]
		//解析token
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		//获取token中的userid
		userId := claims.UserId
		user := mapping.GetUser()
		user, b := user.FindById(userId)
		if !b {
			ctx.JSON(401, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
		}

		//用户存在的话 , 将user信息写入上下文
		ctx.Set("user", user)

		//执行下一个方法
		ctx.Next()
	}
}
