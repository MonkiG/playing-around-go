package main

import "fmt"

func IterateSlice[T any](slice []T) {
	fmt.Println("-----------Iterating-------------")
	for i, v := range slice {
		fmt.Printf("Slice: %v \nIndex: %d \nValue: %v \n", slice, i, v)
	}
}
