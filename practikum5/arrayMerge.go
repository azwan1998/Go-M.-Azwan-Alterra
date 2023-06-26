package main

import (
	"fmt"
)

func arrayMerge(arrayA, arrayB []string) []string {
	mapping := make(map[string]bool)

	for _, elemen := range arrayA {
		mapping[elemen] = true
	}

	for _, elemen := range arrayB {
		if !mapping[elemen] {
			arrayA = append(arrayA, elemen)
			mapping[elemen] = true
		}
	}

	return arrayA
}

func main() {
	arrayA := []string{"kazuya", "jin", "lee"}
	arrayB := []string{"kazuya", "feng"}

	resultArray := arrayMerge(arrayA, arrayB)
	fmt.Println(resultArray)
}
