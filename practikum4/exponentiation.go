package main

import "fmt"

func hitungPangkat(bilangan, pangkat int) int {
	result := 1
	for i := 0; i < pangkat; i++ {
		result *= bilangan
	}
	return result
}

func main() {
	var bilangan, pangkat int
	fmt.Print("Masukkan bilangan: ")
	_, err := fmt.Scanf("%d", &bilangan)
	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	fmt.Print("Masukkan pangkat: ")
	_, err = fmt.Scanf("%d", &pangkat)
	if err != nil {
		fmt.Println("Input tidak valid!")
		return
	}

	result := hitungPangkat(bilangan, pangkat)
	fmt.Printf("%d pangkat %d = %d\n", bilangan, pangkat, result)
}
