package model

type Buku struct {
	Judul     string `json:"judul" `
	Pengarang string `json:"pengarang"`
	Penerbit  string `json:"penerbit"`
}
