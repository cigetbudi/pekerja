package models

import (
	"pekerja/db"
	"time"
)

type Personals struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Phone     string    `json:"phone"`
	Umur      int       `json:"umur"`
	Pekerjaan int       `json:"pekerjaan"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Personals) CreatePersonal() error {
	if err := db.DB.Create(p).Error; err != nil {
		return err
	}
	return nil
}

func GetAll(key string) ([]Personals, error) {
	var ps []Personals
	result := db.DB.Where("nama LIKE?", "%"+key+"%").Find(&ps)
	return ps, result.Error
}

func GetById(id int) (Personals, error) {
	var p Personals
	result := db.DB.Where("id= ?", id).First(&p)
	return p, result.Error
}
