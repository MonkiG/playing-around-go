package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Personable interface {
	Walk()
	eat()
}

func (p *Person) Walk() {
	fmt.Printf("%s is walking\n", p.Name)
}

func (p *Person) eat() {
	fmt.Printf("%s is eating\n", p.Name)
}

type Person struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Age      int       `json:"age"`
}
