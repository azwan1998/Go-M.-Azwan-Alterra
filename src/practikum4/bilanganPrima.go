package main

import "fmt"

func isPrima(bilangan int) bool {
	if bilangan <= 1 {
		return false
	}

	for c := 2; c*c <= bilangan; c++ {
		if bilangan%c == 0 {
			return false
		}
	}

	return true
}

func main() {
	//inputan angka
	var bilangan int
	fmt.Print("Masukkan nilai bilangan: ")
	_, err := fmt.Scanf("%d", &bilangan)
	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	if isPrima(bilangan) {
		fmt.Printf("adalah bilangan prima.\n")
	} else {
		fmt.Printf("bukan bilangan prima.\n")
	}
}
