package user

import (
	"fmt"
	"gin/base"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB


func init() {
	db, _ = gorm.Open("postgres", "host=127.0.0.1 port=5433 user=admin dbname=postgres password=17f16miM1PBio0FX sslmode=disable")
}

func GetUser() (*base.User, error){
	defer db.Close()

	user := base.User{}


	fmt.Println("连接成功")
	err := db.First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return &user, nil
}