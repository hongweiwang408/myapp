package model

import (
	"fmt"
	"myapp/utils"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

)

//声明gorm句柄

var Db *gorm.DB
var err error


func InitDb() {
	dsn := "root:123456" + "@tcp(" + utils.DbHost + ":" + utils.DbPort + ")/" + utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用复数表名
		},
		Logger: logger.Default.LogMode(logger.Silent),
})
	if err != nil {
		fmt.Println(err)
		panic("数据库连接失败")
	}else{
		fmt.Println("数据库连接成功")
}
	err := Db.AutoMigrate(&User{},&Article{})
	if err != nil {
		panic(err)
	}

}

