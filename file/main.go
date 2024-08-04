package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")

	if err != nil{
		fmt.Println("Enter while creatin file: ", err)
		return
	}else{
		defer file.Close()
		fmt.Println("File create successfully")
	}

	content := "hello world by akash"
	_, error := io.WriteString(file, content)

	if error != nil{
		fmt.Println("this is error")
	}
}