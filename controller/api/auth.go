package api

import (
	"WebProject/ginEssential/dto"
	"WebProject/ginEssential/mapping"
	"WebProject/ginEssential/response"
	"WebProject/ginEssential/service/api"
	"WebProject/ginEssential/utils"
	"github.com/gin-gonic/gin"
)

type ApiAuth struct {
}

func (this ApiAuth) Register(ctx *gin.Context) {

	authservice := api.AuthService{}

	//获取参数
	name := ctx.PostForm("name")
	//数据验证
	telephone := ctx.PostForm("telephone")
	//判断手机号是否存在
	pwd := ctx.PostForm("pwd")

	if authservice.RegisterService(ctx, name, telephone, pwd) {
		//返回结果
		response.Success(ctx, gin.H{}, "注册成功")
		//c.JSON(200, gin.H{
		//	"msg": "注册成功",
		//})
	}

}

func (this ApiAuth) Login(ctx *gin.Context) {

	//获取参数
	ctx.PostForm("name")
	//数据验证
	telephone := ctx.PostForm("telephone")
	//判断手机号是否存在
	pwd := ctx.PostForm("pwd")

	user := api.LoginService(telephone, pwd)
	if len(user.Id) == 0 {
		response.Fail(ctx, gin.H{}, "登录失败")
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//	"code": "500",
		//	"msg":  "登录失败",
		//})
		return
	}

	token, err := utils.ReleaseToken(*user)
	if err != nil {
		response.Fail(ctx, gin.H{}, "签发token失败")
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//"code": "500",
		//"msg":  "签发token失败",
		//})
		return
	}

	response.Success(ctx, gin.H{"token": token}, "登录成功")
	//c.JSON(http.StatusOK, gin.H{
	//	"code": "200",
	//	"data": gin.H{"token": token},
	//	"msg":  "登录成功",
	//})
	//"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMCIsImV4cCI6MTY3MDA3MjU3OSwiaWF0IjoxNjY5NDY3Nzc5LCJpc3MiOiJvY2VhbmxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.PjoG7Nzu4H4IGNS_CQg6zgxQo-rGna3Kzru3LKZKg6o"
	return

}

// 获取用户信息   之前需要有token
func (this ApiAuth) UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(*user.(*mapping.User))}, "获取成功")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "获取成功",
	//	"data": gin.H{
	//		"user": dto.ToUserDto(*user.(*mapping.User)),
	//	},
	//})
	return

}
