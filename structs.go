package main

import "fmt"

type User struct{
	name string
	email string
	age int
}

func main() {
	// Declaring and initializing anonymous struct
	anonymous := struct {
		name  string
		age   int
		email string
	}{
		name:  "X",
		email: "x@gmai.com",
		age:   24,
	}
	fmt.Println("Anonymous Details")
	// Display the field values
	fmt.Println("Name : ", anonymous.name)
	fmt.Println("Email : ", anonymous.email)
	fmt.Println("Age : ", anonymous.age)

	// Named struct

	// Declare variable of type user and init using a struct literal

	joe := User{
		name : "Joe Orton",
		email : "joe@gmail.com",
		age : 24,
	}
	fmt.Println("Joe Details")
	// Display the field values
	fmt.Println("Name : ", joe.name)
	fmt.Println("Email : ", joe.email)
	fmt.Println("Age : ", joe.age)
}
