package model

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"myapp/utils/errmsg"
)


type User struct {
	gorm.Model `json:"gorm.Model"`
	Email      string `gorm:"type:varchar(30);not null" json:"email" `
	Name       string `gorm:"type:varchar(20);not null" json:"name" .validate:"required,min=4,max=12"`
	Password   string `gorm:"type:varchar(20);not null" json:"password" .validate:"required,min=6 max=20"`
	Status     int    `gorm:"type:int;DEFAULT:2" json:"status" .validate:"required,gte=2"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user User
	//在id列表中查询username=name的第一个数据
	Db.Select("id").Where("name = ?", name).First(&user)
	fmt.Println(user.ID)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["status"] = data.Status
	err := Db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err := Db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// BeforeSave 回调函数(钩子函数)//不用被调用
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)
	return nil
}

// ScryptPassword 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	// todo
	salt := make([]byte, 8)
	salt = []byte{3, 44, 6, 7, 32, 45, 99, 0}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

//登录验证

func CheckLogin(username string, password string) int {
	var user User
	Db.Where("name=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ERROR_PASSWORD_ERROR
	}
	if user.Status == 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
