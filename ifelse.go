package main

import (
	"fmt"
)

func main(){
	// check eligible for bote or not
	age := 20

	if (age >= 18){
		fmt.Print("yes you are eligible for bote! you can bote")
	} else{
		fmt.Print("NO you are not elegible for bote!")
	}
}