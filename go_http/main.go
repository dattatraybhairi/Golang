package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	for i := 0; i < len(array); i++ {
		fmt.Println(&array[i])
	}
	array1 := &array
	for i := 0; i < len(array1); i++ {
		fmt.Println(&array1[i])
	}
}
