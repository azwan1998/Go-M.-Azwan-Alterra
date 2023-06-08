package main

import (
	"fmt"
)

func hitungLuasPermukaanTabung(jariJari, tinggi float64) float64 {
	pi := 3.14
	luasAlas := pi * jariJari * jariJari
	luasSelimut := 2 * pi * jariJari * tinggi
	luasPermukaan := 2*luasAlas + luasSelimut

	return luasPermukaan
}

func main() {
	T := 20.0
	r := 4.0
	luasPermukaan := hitungLuasPermukaanTabung(T, r)
	fmt.Println("Luas permukaan tabung adalah:", luasPermukaan)
}
