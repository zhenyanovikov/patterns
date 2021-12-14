package main

import (
	"fmt"
)

type SimpleFabricPattern struct {
}

func (p SimpleFabricPattern) Show() {
	printAnimalSound(animalFabric("cat"))
	printAnimalSound(animalFabric("dog"))
}

func (p SimpleFabricPattern) Name() string {
	return "SimpleFabric"
}

func printAnimalSound(a Animal) {
	fmt.Printf("I'm %s and i say %s\n", a.AnimalName(), a.Sound())
}

func animalFabric(name string) Animal {
	switch name {
	case "cat":
		return Cat{}
	case "dog":
		return Dog{}
	default:
		return nil
	}
}

type Animal interface {
	AnimalName() string
	Sound() string
}

type Cat struct {
}

func (c Cat) AnimalName() string {
	return "Cat"
}

func (c Cat) Sound() string {
	return "Meow"
}

func (c Cat) String() string {
	return fmt.Sprintf("I'm %s and i say %s", c.AnimalName(), c.Sound())
}

type Dog struct {
}

func (c Dog) AnimalName() string {
	return "Dog"
}

func (c Dog) Sound() string {
	return "Bow-wow"
}
