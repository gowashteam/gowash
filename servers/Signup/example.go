package main

import "fmt"

type person struct {
	name    string
	address string
	school  string
}

//method to return the person object
func getperson() *person {
	aPerson := person{} // an empty person object
	aPerson.name = "Okey"
	name := aPerson.name
	fmt.Printf("My name is %v \n", name)

	return &aPerson
}

func main2() {

	name := getperson().name
	fmt.Printf("I can also appear here as %v", name)
}
