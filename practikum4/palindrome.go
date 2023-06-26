package main

import (
	"fmt"
)

func isPalindrome(input string) bool {

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var kata string
	fmt.Print("Masukkan satu kata:")
	_, err := fmt.Scanf("%s", &kata)
	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	if isPalindrome(kata) {
		fmt.Println(kata, "adalah palindrome.")
	} else {
		fmt.Println(kata, "bukan palindrome.")
	}
}
