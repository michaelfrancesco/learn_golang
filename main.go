package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var firstName string = "Michael"
	var lastName string
	lastName = "Francesco"

	// type inference
	umurSaya := "jonathan"

	// Multiple variable
	var anak1, anak2, anak3 string
	anak1, anak2, anak3 = "jason", "juliet", "john"
	umur1, umur2, umur3 := 10, 9, 7

	// variable _

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("Umur: ", umurSaya)
	fmt.Println(anak1, anak2, anak3)
	fmt.Println("Umur anak:", umur1, umur2, umur3)

}
