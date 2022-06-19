package models

import "time"

type Personals struct {
	Id        int
	Nama      string
	Alamat    string
	Phone     string
	Umur      int
	Pekerjaan int
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
