package api

import (
	"WebProject/ginEssential/mapping"
	"WebProject/ginEssential/response"
	"WebProject/ginEssential/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type AuthService struct {
}

// 验证
func (this *AuthService) RegisterService(ctx *gin.Context, name, telephone, pwd string) bool {

	//数据验证
	if len(telephone) != 11 {
		response.Fail(ctx, gin.H{}, "手机号不正确")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"msg":  "手机号不正确",
		//	"code": 422,
		//})
		ctx.Abort()
		return false
	}

	if len(pwd) < 6 {
		response.Fail(ctx, gin.H{}, "密码小于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "密码小于6位",
		//})
		ctx.Abort()
		return false
	}

	if len(name) == 0 {
		name = string(utils.Random(10))
	}

	var user = mapping.GetUser()
	//验证手机号
	u := user.TelephonIsexist(telephone)
	if len(u.Telephone) != 0 {
		response.Fail(ctx, gin.H{}, "手机号已存在")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号已存在",
		//})
		ctx.Abort()
		return false
	}

	b := user.AddUser(name, pwd, telephone)

	if !b {
		response.Fail(ctx, gin.H{}, "添加角色错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{
		//	"code": 500,
		//	"msg":  "添加角色错误",
		//})
		ctx.Abort()
		return false
	}

	return true
}

func LoginService(telephone, pwd string) *mapping.User {

	var user = mapping.GetUser()
	user = user.Login(pwd, telephone)
	if len(user.Id) == 0 {
		log.Println("登陆失败")
		return &mapping.User{}
	}
	return user
}
