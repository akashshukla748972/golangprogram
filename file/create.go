package main

import (
	"fmt"
	"os"
)

func createFile() {
	file, err := os.Create("first.txt")

	if err != nil{
		fmt.Printf("creating file geting some error %s",err)
	}else{
		fmt.Printf("file created successfully %v",file)
	}
	defer file.Close()
}