package controllers

import (
	"net/http"
	"pekerja/helpers"
	"pekerja/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	if !res {
		return echo.ErrUnauthorized
	}
	// return c.String(http.StatusOK, "berhasil login")
	// generate token

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login Berhasil!",
		"token":   t,
	})
}

func GenerateHashPass(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func Register(c echo.Context) error {
	u := new(models.Users)
	c.Bind(u)

	r := new(models.Response)
	if u.CreateUser() != nil { //method CreateUser
		r.ErrorCode = 10
		r.Message = u.CreateUser().Error()
	} else {
		r.ErrorCode = 0
		r.Message = "Berhasil Registrasi"
		r.Data = *u
	}
	return c.JSON(http.StatusCreated, r)
}
