package base

import "github.com/jinzhu/gorm"

type User struct {
	Id     int    `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
}

type OrmDb struct {
	db *gorm.DB
}