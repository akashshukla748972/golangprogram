package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Printf("current Date and Time is %v\n", currentTime)

	date := currentTime.Format("02-01-2006 January")
	fmt.Printf("current date is :%v\n", date)

	time := currentTime.Format("3:04 PM")
	fmt.Printf("current date is :%v\n", time)


}