package v1

import (
	"github.com/gin-gonic/gin"
	"myapp/model"
	"myapp/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加帖子
func AddArt(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)
	if code == errmsg.SUCCESS {
		model.Db.Preload("User").First(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArt  查询帖子
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageNum == 0 {
		pageNum = -1
	}
	data, code := model.GetArt(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除帖子
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //获取id
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
