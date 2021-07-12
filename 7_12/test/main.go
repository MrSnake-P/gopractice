package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	//gorm.Model
	ID		int64
	Code 	string
	Price 	string
}
type User struct {
	Id     int    `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
}


func main() {
		db, _ := gorm.Open("postgres", "host=127.0.0.1 port=5433 user=admin dbname=postgres password=17f16miM1PBio0FX sslmode=disable")
		defer db.Close()
		//user := User{}
		a := db.Find(&[]Product{})
		//result := db.Where("id = ?",1)
		//result.Last(&user)
		fmt.Println(a.Value)
		//fmt.Println("连接成功")
		//err := db.First(&user).Error
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(user)
}
