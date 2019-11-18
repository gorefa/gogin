package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

func main() {
	db, _ := gorm.Open("mysql", "root:NuCkQlqckmp@tcp(test.app.ifengidc.com:58998)/gogin?charset=utf8&parseTime=True&loc=Local")

	//var user = User{Age: 26, Name: "lisai"}

	if !db.HasTable(&User{}) {
		fmt.Println("table users not exist!")
		db.CreateTable(&User{})
	}

	//db.Create(&user)

	user := []User{}
	db.Where("name = ?", "lisai").Find(&user)
	for _, v := range user {
		fmt.Println("name: age:", v.Name, v.Age)
	}

}
