package main

import (
	"myapp/model"
	"myapp/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
		// 	model.CreateUser(&model.User{
		// 	Email: "www@qq.com",
		// 		Name: "hw",
		// 		Password:"123456",
		// })
	// fmt.Println(model.CheckLogin("hw", "123456"))
	// fmt.Println(model.CheckUser("hw"))
	// fmt.Println(model.GetUsers(10,1))

	// fmt.Println(model.DeleteUser(3))
}
