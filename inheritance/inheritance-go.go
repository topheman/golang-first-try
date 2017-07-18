package main

import (
	"fmt"
)

// Animal ...
type Animal struct {
	name string
}

func (animal *Animal) eat(foodName string) {
	fmt.Println(animal.name + " is eating " + foodName)
}

// Dog ...
type Dog struct {
	Animal
	isPet bool
}

func (dog *Dog) eat() {
	if dog.isPet {
		dog.Animal.eat("dog food")
		return
	}
	fmt.Println(dog.Animal.name + " is happy to eat!")
}

func (dog *Dog) bark() {
	fmt.Println(dog.Animal.name + " barks!")
}

func main() {
	myDog := Dog{Animal{"Georges"}, true}
	myDog.eat()
	myDog.bark()
	regularDog := Dog{Animal{"Jules"}, false}
	regularDog.eat()
	regularDog.bark()
}
