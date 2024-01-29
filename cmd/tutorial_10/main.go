package main

import "fmt"
func main (){
	var c = make(chan int) //channel created 
	go process(c) // start goroutine
	for i:= range c{ //waiting..
	
	fmt.Println(i)
}}
func process(c chan int){
	defer close(c) //i
	for i:=0;i<5;i++{ //setup for loop
		c <- i
	}
	//close(c) ii
	//either i or ii works 
}