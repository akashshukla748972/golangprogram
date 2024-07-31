package main

import (
	"fmt"
)

func main(){
	fmt.Println("Print table")

	var n int = 5
	for i:= 1; i<=10; i++{
		fmt.Println(i*n)
	}
}