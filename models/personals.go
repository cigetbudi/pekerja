package models

import "time"

type Personals struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Phone     string    `json:"phone"`
	Umur      int       `json:"umur"`
	Pekerjaan int       `json:"pekerjaan"`
	IsActive  bool      `json:"isactive"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
