package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	slice := []int{1, 2, 1, 5, 8, 2}
	fmt.Println(slice)
}
