package model

import (
	"fmt"

	"gogin/pkg/auth"

	"gopkg.in/go-playground/validator.v9"
)

var DefaultLimit = 50

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Create() error {
	return DB.Create(&u).Error
}

func ListUser() {
	fmt.Println("11111")
	users := []UserModel{}
	err := DB.Select("username").Find(&users).Error
	if err != nil {
		fmt.Println("err", err)
	}
	for _, v := range users {
		fmt.Println("nnnnnnn", v)
	}
}
