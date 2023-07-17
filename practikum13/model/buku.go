package model

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	Judul     string `json:"judul" `
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit"`
}
