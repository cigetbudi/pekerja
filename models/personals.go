package models

import (
	"pekerja/db"
	"time"
)

type Personals struct {
	Email     string    `json:"email" gorm:"primaryKey"`
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

func (p *Personals) UpdatePersonal(email string) error {
	if err := db.DB.Model(&Personals{}).Where("email =?", email).Updates(p).Error; err != nil {
		return err
	}
	return nil
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
