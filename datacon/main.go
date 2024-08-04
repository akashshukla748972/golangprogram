package main

import (
	"fmt"
	"strconv"
	// "strconv"
)

func main(){
	var num int = 42
	fmt.Printf("Number is %v\n", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("data is %v\n", data)
	fmt.Printf("Type of data is %T\n", data)

	str := strconv.Itoa(num) + "Akash"
	fmt.Printf("str is %v\n", str)
	fmt.Printf("type of str is %T\n", str)

	number := "12345"
	str_int, _:= strconv.Atoi(number)
	fmt.Printf("str_int is %v\n", str_int)
	fmt.Printf("type of str_int is %T\n", str_int)


}