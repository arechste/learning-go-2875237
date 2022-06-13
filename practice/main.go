package main

import (
	"fmt"
)

func main() {
	fmt.Println("func custom types")
	poodle := Dog{"Poodle", 10, "Wuff"}
	fmt.Println(poodle)
	// print field/value of struct
	fmt.Printf("%+v\n", poodle)
	// print fields of struct
	fmt.Printf("Breed: %v\nWeight %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

//Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

//Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

//SpeakThreeTimes is multiSpeak dog speak
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
