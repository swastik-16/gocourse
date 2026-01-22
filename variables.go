package main

import "fmt"

var middleName string = "Not applicable"
//:= can only be used in a function

func main() {
	var age int
	var name string = "Swastik"
	var name1 string = "Raj"
	count := 10
	lastName := "Mukherjee"
	fmt.Println(middleName)
	fmt.Println(age,name,name1,count,lastName)
	//Default values
	//Numeric - 0
	//Boolean - false
	//String - ""
	//Pointers,slices,maps,structs,functions - nil

	// --- SCOPE
	//variables have block scope
	printName()
}

func printName() {
	firstName := "Swastik"
	fmt.Println(firstName)
}