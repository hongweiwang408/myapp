package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"username"`
	Pwd  string `json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/json",bindJson)
	r.POST("/json2",bindJson2)
	r.Run()
	fmt.Println("hello,world")
}


func bindJson(c *gin.Context){
	var user User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		log.Println(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":"its ok",
	})

}

func bindJson2(c *gin.Context)  {
	var user User
	// first call ShouldBindJson
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
	}
	log.Printf("username: %s, password: %s", user.Name,user.Pwd)
	// second call ShouldBindJson
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // EOF
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "bind json is ok"})
}