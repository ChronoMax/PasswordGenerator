package main

import (
	"fmt"
	"strconv"

	"github.com/ChronoMax/PasswordGenerator/generation"
	"github.com/ChronoMax/PasswordGenerator/mysqlscript"
)

func main() {
	var input string

	//Displays text and gets the users input.
	fmt.Println("How long should the password be? min 15 max 20 characters.")
	fmt.Scanln(&input)

	//Converts the string input to a int.
	convertedAmount, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Could not convert the users input.")
	}

	if convertedAmount <= 14 || convertedAmount >= 21 {

		//Give output
		fmt.Println("Warning! Please note the characterl limit!")

	} else {
		//Displays the returned value.
		b := generation.GeneratePassword(convertedAmount)
		mysqlscript.InsertToDatabase(b)
		fmt.Println(b)
	}
}
