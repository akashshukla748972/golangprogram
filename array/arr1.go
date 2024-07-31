package main

import "fmt"

func main() {
	var name [5]string

	name[0] = "Akash"
	name[1] = "Vikash"
	name[2] = "Anmol"
	name[4] = "Pranshu"

	var age = [5]int{24, 22, 18, 20}

	str := "hello world"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Print(name[0][1:4])
	fmt.Print(str[1:4])
}