package main

import "fmt"

func printSum(a,b int) (int){
	return a+b
}

func main() {
	fmt.Println("Starting of the program")
	data := printSum(12, 15)
	defer fmt.Printf("sum is %v\n", data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}