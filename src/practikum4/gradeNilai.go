package main

import (
	"fmt"
)

func gradeNilai(score int) string {
	if score >= 80 && score <= 100 {
		return "A"
	} else if score >= 65 && score <= 79 {
		return "B"
	} else if score >= 64 && score <= 50 {
		return "C"
	} else if score >= 49 && score <= 35 {
		return "D"
	} else if score >= 0 && score <= 34 {
		return "E"
	} else {
		return "Nilai Invalid"
	}

}

func main() {

	var studentScore int
	fmt.Print("Masukkan nilai siswa: ")
	_, err := fmt.Scanf("%d", &studentScore)
	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	grade := gradeNilai(studentScore)
	fmt.Println("nilai siswa:", studentScore)
	fmt.Println("Grade nilai siswa ini adalah:\n", grade)
}
