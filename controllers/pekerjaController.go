package controllers

import (
	"net/http"
	"pekerja/models"

	"github.com/labstack/echo/v4"
)

func FetchAllPersonal(c echo.Context) error {
	response := new(models.Response)
	ps, err := models.GetAll(c.QueryParam("nama")) //method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal mengambil data pekerja"
	} else if len(ps) == 0 {
		response.ErrorCode = 10
		response.Message = "Pekerja tidak dapat ditemukan"
		response.Data = ps
	} else {
		response.ErrorCode = 0
		response.Message = "Pekerja ditemukan"
		response.Data = ps
	}
	return c.
		JSON(http.StatusOK, response)
}

func FetchPersonal(c echo.Context) error {
	response := new(models.Response)
	ps, err := models.GetAll(c.QueryParam("nama")) //method get all
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal mengambil data pekerja"
	} else if len(ps) == 0 {
		response.ErrorCode = 10
		response.Message = "Pekerja tidak dapat ditemukan"
		response.Data = ps
	} else {
		response.ErrorCode = 0
		response.Message = "Pekerja ditemukan"
		response.Data = ps
	}
	return c.
		JSON(http.StatusOK, response)
}

func AddPersonal(c echo.Context) error {
	p := new(models.Personals)
	c.Bind(p)

	response := new(models.Response)
	if p.CreatePersonal() != nil { //method create
		response.ErrorCode = 10
		response.Message = "Gagal menambahkan data personal"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses menambahkan data user"
		response.Data = *p
	}
	return c.JSON(http.StatusOK, response)
}

func EditPersonal(c echo.Context) error {
	p := new(models.Personals)
	c.Bind(p)
	response := new(models.Response)
	if p.UpdatePersonal(c.Param("id")) != nil { //method update
		response.ErrorCode = 10
		response.Message = "Gagal dalam mengedit data personal"
	} else {
		response.ErrorCode = 0
		response.Message = "Berhasil edit data personal"
		response.Data = *p
	}
	return c.JSON(http.StatusOK, response)
}
