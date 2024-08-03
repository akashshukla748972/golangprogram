package main

import "fmt"

type Student struct{
	sid int
	name string
	course string
	fees int
}

func main() {
	fmt.Print("this is struct")
	var student1 Student
	student1.sid = 101
	student1.name = "Akash kumar shukla"
	student1.course = "BCA"
	student1.fees = 16000

	fmt.Printf("id %d and name of the student %s course is %s and fees %d", student1.sid, student1.name, student1.course, student1.fees)
	 
}