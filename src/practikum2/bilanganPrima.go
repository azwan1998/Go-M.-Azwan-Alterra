package main

import "fmt"

func isPrima(number int) bool {
	if number <= 1 {
		return false
	}

	for c := 2; c*c <= number; c++ {
		if number%c == 0 {
			return false
		}
	}

	return true
}

func main() {
	//inputan angka
	number := 2

	if isPrima(number) {
		fmt.Printf("%d adalah bilangan prima.\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", number)
	}
}
