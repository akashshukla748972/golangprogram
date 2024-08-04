package main

import (
	"fmt"
)

func main(){
	fmt.Println("Map date type")
	var student = map[string]string{
		"name": "Akash",
		"course": "BCA",
		"fees": "16000"	}

		fmt.Printf("student data : %v\n",student)
		fmt.Println("student data :",student)
	
	var employee = make(map[string]string)
	employee["name"] = "laber"
	employee["dep"] = "laber"
	employee["salary"] = "4000"

	fmt.Printf("date of employee is:%v",employee)
}