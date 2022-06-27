package models

import (
	"fmt"
	"pekerja/db"
	"pekerja/helpers"
)

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckLogin(un, pass string) (bool, error) {
	// var pwd string
	var u Users
	result := db.DB.Where("username = ?", un).First(&u)

	if result == nil {
		fmt.Println("User tidak ada")
	}

	pwd := u.Password
	match, err := helpers.CheckPasswordHash(pass, pwd)

	if !match {
		fmt.Println("Hash tidak cocok")
		return false, err
	}
	return true, nil
}

func (u *Users) CreateUser() error {
	var res error
	hPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		res = err
	}
	u.Password = hPass
	if err := db.DB.Create(u).Error; err != nil {
		res = err
	}
	return res
}
