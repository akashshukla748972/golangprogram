package main

import (
	"fmt"
)

func simplefunction(){
	fmt.Println("simple function")
}

func addfunction(a, b int) (int){
	return a+b
}

func main(){
	fmt.Println("we are lerning function in Golang!")
	simplefunction()

	output := addfunction(12, 13)
	fmt.Printf("Sum is :%v",output)
}