package models

import (
	"pekerja/db"
	"time"

	"github.com/go-playground/validator/v10"
)

type Personals struct {
	Email     string    `json:"email" gorm:"primaryKey" validate:"required,email"`
	Nama      string    `json:"nama" validate:"required"`
	Alamat    string    `json:"alamat"`
	Phone     string    `json:"phone" validate:"e164"`
	Umur      int       `json:"umur" validate:"gte=0,lte=130"`
	Pekerjaan int       `json:"pekerjaan" validate:"required,numeric"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Personals) CreatePersonal() error {
	v := validator.New()
	var res error
	if err := v.Struct(p); err != nil {
		res = err
	} else if err := db.DB.Create(p).Error; err != nil {
		res = err
	}
	return res
}

func (p *Personals) UpdatePersonal(email string) error {
	v := validator.New()
	var res error
	if err := v.Struct(p); err != nil {
		res = err
	} else if err := db.DB.Model(&Personals{}).Where("email =?", email).Updates(p).Error; err != nil {
		res = err
	}
	return res
}

func (p *Personals) DeletePersonal() error {
	if err := db.DB.Delete(p).Error; err != nil {
		return err
	}
	return nil
}

func GetAll(key string) ([]Personals, error) {
	var ps []Personals
	result := db.DB.Where("nama LIKE?", "%"+key+"%").Find(&ps)
	return ps, result.Error
}

func GetByEmail(email string) (Personals, error) {
	var p Personals
	result := db.DB.Where("email= ?", email).First(&p)
	return p, result.Error
}
