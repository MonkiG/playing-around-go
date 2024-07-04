package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello")

	ramon := Person{
		Id:       uuid.New(),
		Name:     "Ramón",
		Email:    "some@gmail",
		Password: "somepassword",
		Age:      25,
	}

	ramon.eat()
	ramon.Walk()
	isEven, err := IsEven(ramon.Age)

	if err != nil {
		fmt.Println(err)
	}

	if isEven {
		fmt.Printf("The age of %s is even\n", ramon.Name)
	} else {
		fmt.Printf("The age of %s is not even\n", ramon.Name)
	}

	personD, _ := json.Marshal(ramon)
	fmt.Printf("Mi person: %+v\n", string(personD))

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	IterateSlice(arr[:])
	yieldNumber := IterateNumber(0)

	for {
		if i := yieldNumber.Next(); i <= 10 {
			fmt.Println(i)
		} else {
			break
		}
	}

	Sync()
	Conc()
	fmt.Println(Friday)
	time.Sleep(time.Second)
}
