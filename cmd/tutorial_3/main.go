package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 8
	var denominator int = 3
	result, remainder, err := intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0  {
	// 		fmt.Printf("The result of the integer division is %v \n", result)
		
	// } else{
	// 	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	// }
	switch{
	case err!=nil: 
	fmt.Println(err.Error())
	case remainder == 0 : 
			fmt.Printf("The result of the integer division is %v \n", result)	
			default:
				fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)		
	}
}

func intDivision(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return
	}
	result = numerator / denominator
	remainder = numerator % denominator
	return result, remainder, nil
}
