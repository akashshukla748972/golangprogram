package main

import "fmt"

func modifyValueByReference(num *int){
	*num = *num + 20
}

func main() {
	// var a int = 2
	// fmt.Printf("value of a is %d", a)

	// var ptr *int
	// ptr = &a
	// fmt.Printf("value of pointer contain %v",ptr)

	num := 2
	ptr := &num

	fmt.Printf("valu of num %d and poiner %v contain\n",num,ptr)//print pointer container address
	fmt.Printf("valu of num %d and poiner %v contain\n",num,*ptr)//print witch value pointer containe address

	//modify value using pointer
	modifyValueByReference(&num)
	fmt.Printf("valu of num %d and poiner %v contain\n",num,ptr)
}