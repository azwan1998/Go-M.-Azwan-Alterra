package main

import "fmt"

func getMinMax(numbers []int) (int, int) {
	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	numbers := make([]int, 6)

	fmt.Println("Enter 6 numbers:")

	for i := 0; i < 6; i++ {
		fmt.Scan(&numbers[i])
	}

	max, min := getMinMax(numbers)

	fmt.Printf("%d is the maximum number\n", max)
	fmt.Printf("%d is the minimum number\n", min)
}
