package main

import (
	"fmt"
)

func main() {

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}

// struct exported Dog / private dog
//Dog is a struct
type Dog struct {
	//exported fields
	Breed  string
	Weight int
}
