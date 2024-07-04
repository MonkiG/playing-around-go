package main

import (
	"fmt"
	"strconv"
)

func say(message string) {
	fmt.Println(message)
}

func loop(limit int, callback func(i int)) {
	for i := 0; i <= limit; i++ {
		callback(i)
	}
}

func Sync() {
	fmt.Println("-----Sync loop-----")
	loop(10, func(i int) {
		message := "Hello " + strconv.Itoa(i)
		say(message)
	})
}

func Conc() {
	fmt.Println("-----Concurrently loop-----")
	loop(10, func(i int) {
		message := "Hello " + strconv.Itoa(i)
		go say(message)
	})
}
