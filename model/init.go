package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func Init() {
	var err error
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		true,
		"Local")

	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", viper.GetString("db.username"))
	}

	DB.DB().SetMaxOpenConns(2000)
	DB.DB().SetMaxIdleConns(0)

	//defer DB.Close()

}
