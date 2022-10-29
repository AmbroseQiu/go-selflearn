package pack1

import (
	"fmt"
)

type Contact struct {
	ID          string
	PhoneNumber string
	Gender      int
}

type Animal interface {
	bark()
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func (cat_ *Cat) setName() {
	cat_.name = "cat"
}

func (dog_ *Dog) setName() {
	dog_.name = "dog"
}

func (cat_ *Cat) bark() {
	fmt.Println(cat_.name, " meow")
}

func (dog_ *Dog) bark() {
	fmt.Println(dog_.name, " wan")
}

func animalBark(animal_ Animal) {
	switch animal := animal_.(type) {
	case *Cat:
		animal.setName()
	case *Dog:
		animal.setName()
	}
	animal_.bark()
}

func CallFunc() {
	contact := &Contact{
		ID:          "Ambrose",
		PhoneNumber: "1231",
		Gender:      1,
	}

	fmt.Printf("ID=%v\n", contact.ID)
	animal := Cat{}
	animalBark(&animal)
}
