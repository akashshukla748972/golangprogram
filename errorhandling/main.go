package main

import (
	"fmt"
)

func divide(a, b float32)(float32, error){
	if b == 0{
		return 0, fmt.Errorf("denominator must not be Zero")
	}
	return a/b, nil
}

func main(){
	fmt.Println("started error handling in the function!")
	output, _:= divide(10, 0)
	fmt.Printf("divide value is :%v",output)
}