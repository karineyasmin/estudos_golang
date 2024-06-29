package main

import "fmt"

func main() {
	var number float32

	fmt.Println("Insert a float number: ")
	fmt.Scan(&number)

	truncateNumber := int(number)

	fmt.Println("truncate number:", truncateNumber)
}
