package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "akash,vikash,anmol,ayush,pranshu,anshu"
	fmt.Println(str)

	parts := strings.Split(str, ",")
	fmt.Println(parts)

	str1 := strings.Join(parts, "-")
	fmt.Println(str1)

	str2 := "AkaSh"
	str3 := strings.ToLower(str2)
	fmt.Print(str2)
	fmt.Print(str3)
}