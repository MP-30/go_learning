package main

import "fmt"

var (
	name, location string
	age            int
)

var (
	gender, father_name, mother_name string = "male", "mr.", "mrs."
)

func main() {
	name = "Aditya"
	location = "India"
	age = 25

	fmt.Println(name)
	fmt.Println(location)
	fmt.Println(age)
}
