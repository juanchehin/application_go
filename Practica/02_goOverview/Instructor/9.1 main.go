package main

import (
	"fmt"
)

// Animal defines the interface for type Animal
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Dog defines the dog type
type Dog struct {
	Name  string
	Breed string
}

// Gorilla defines the gorilla type
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	// Create a dog variable of type Dog.
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	fmt.Println(fmt.Sprintf("The dog's name is %s and he is a %s.", dog.Name, dog.Breed))

	// Create a gorilla variable of type Gorilla.
	gorilla := Gorilla{
		Name:          "Geraldine",
		Color:         "black",
		NumberOfTeeth: 32,
	}
	fmt.Println(fmt.Sprintf("The gorilla's name is %s. She has %d teeth and she is %s in color.",
		gorilla.Name, gorilla.NumberOfTeeth, gorilla.Color))

	// We can pass dog to Riddle(), since the Dog type implements Animal interface by having all of the
	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
	// https://tour.golang.org/methods/8 for more details).
	Riddle(&dog)

	// We can also pass gorilla to Riddle(), since the Gorilla type satisfies the Animal interface.
	Riddle(&gorilla)
}

// Says has a receiver of type *Dog, so it satisfies part of the interface requirements for Animal
// for the Dog type
func (d *Dog) Says() string {
	return "woof"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Dog type
func (d *Dog) NumberOfLegs() int {
	return 4
}

// Says has a receiver of type *Gorilla, so it satisfies part of the interface requirements for Animal
// for the Gorilla type
func (g *Gorilla) Says() string {
	return "grunt"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Gorilla type
func (g *Gorilla) NumberOfLegs() int {
	return 2
}

// Riddle asks a riddle
func Riddle(a Animal) {
	info := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is this?`, a.Says(), a.NumberOfLegs())
	fmt.Println(info)
}
