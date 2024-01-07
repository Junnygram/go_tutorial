package main

import (
	"fmt"
	"strings"
)
func main (){
	myString := []rune("résume")
	indexed := myString[0]
	fmt.Println(myString)
	fmt.Println(indexed) //114
	fmt.Printf("%v, %T\n", indexed, indexed ) // this prints the ascii character and the type =  114 , uint8
	for i, v := range myString{
		fmt.Println(i, v)
	}

	// Strings are immutable in go meaning it cannot be modified once created 
	strSlice := []string{"O", "l", "a", "l", "e", "y", "e"}
	catStr := ""
	for i:= range strSlice{
		catStr += strSlice[i]
	}
	fmt.Printf("\n %v \n", catStr) // olaleye

	//or

	var strBuilder strings.Builder
	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	newStr := strBuilder.String()
	fmt.Printf("\n %v \n", newStr) // olaleye
}