package response

import (
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, data gin.H, msg string) {

	ctx.JSON(httpStatus, gin.H{
		"code": httpStatus,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, 400, data, msg)
}
