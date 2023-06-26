package main

import (
	"fmt"
)

func munculSekali(input string) []int {
	mapping := make(map[rune]int)

	for _, char := range input {
		mapping[char]++
	}

	var angkaUnik []int
	for _, char := range input {
		if mapping[char] == 1 {
			angkaUnik = append(angkaUnik, int(char-'0'))
		}
	}

	return angkaUnik
}

func main() {
	input := "76523752"
	hasil := munculSekali(input)
	fmt.Println(hasil)

	input = "1122"
	hasil = munculSekali(input)
	fmt.Println(hasil)
}
