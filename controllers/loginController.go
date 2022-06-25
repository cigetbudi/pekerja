package controllers

import (
	"net/http"
	"pekerja/helpers"
	"pekerja/models"

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
	return c.String(http.StatusOK, "berhasil login")
	//generate token

}

func GenerateHashPass(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}