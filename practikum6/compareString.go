package main

import (
	"fmt"
	"strings"
)

func compare(A, B string) string {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j <= len(A); j++ {
			substring := A[i:j]
			if strings.Contains(B, substring) {
				return substring
			}
		}
	}
	return ""
}

func main() {
	A := "AKA"
	B := "AKASHI"
	commonSubstring := compare(A, B)
	fmt.Println(commonSubstring)

	A = "KANGAOORO"
	B = "KANG"
	commonSubstring = compare(A, B)
	fmt.Println(commonSubstring)
}
