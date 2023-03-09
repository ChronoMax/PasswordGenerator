package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("How long should the password be? min 15 max 20 characters.")
	tempAmount := fmt.Scanln()
	convertedAmount, err := strconv.Atoi(tempAmount)

	GeneratePassword(convertedAmount)
	fmt.Println(b)
}
