package v1

import (
	"myapp/middleware"
	"myapp/model"
	"myapp/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var code int
	var token string
	code = model.CheckLogin(data.Name, data.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
