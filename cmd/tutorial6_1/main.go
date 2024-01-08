package main

import "fmt"
type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
}
type owner struct{
	name string
}

func main(){
	var myEngine gasEngine = gasEngine{25, 15, owner{"Ola"}}
	var myEngine2  = struct{
		mpg uint8
		gallons uint8
		name string
	}{35,30, "Junny"}
	//differences gasEngine is reusable and gasEngine2 is not reusable
	fmt.Println(myEngine.mpg ,myEngine.gallons, myEngine.ownerInfo.name ) // 25 15 Ola
	fmt.Println(myEngine2.mpg ,myEngine2.gallons, myEngine2.name )  //35 30 Junny
}
