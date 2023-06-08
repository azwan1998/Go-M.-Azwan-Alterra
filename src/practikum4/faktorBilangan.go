package main

import (
	"fmt"
)

func faktorBilangan(bilangan int) {
	fmt.Printf("Faktor dari bilangan %d adalah: \n", bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println()

}

func main() {

	var bilangan int
	fmt.Print("Masukkan nilai bilangan: ")
	_, err := fmt.Scanf("%d", &bilangan)
	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	faktorBilangan(bilangan)

}
