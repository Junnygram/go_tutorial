package main

import "fmt"
func main(){
var p *int32 = new(int32) //here we tell the compiler we want to initialise a pointer
var i int32
fmt.Printf("The value p point is :%v", *p) // here we tell the compiler we want to refrence the value of the pointer
fmt.Printf("\nThe value of i is: %v\n", i)
*p = 10 // here we tell the compiler we want to refrence the value of the pointer
 p= &i//memory address of the variable not the value

//example below
var slice = []int32{1,2,3}
var sliceCopy = slice
sliceCopy[2] = 4
fmt.Println(slice) //[1 2 4]
fmt.Println(sliceCopy) //[1 2 4]


//thing1 and thing2 have different memory location meaning they are differnet array so we can modify the value of thing2 without affecting the value of thing1
//we use pointer to give the same memory location
//now when pointer is used changing the value of thing1 affects thing2
var thing1 = [5]float64 {1,2,3,4,5}
fmt.Printf("\n The memory location of the thing1 array is: %p", &thing1)
var result [5]float64 = square(&thing1)
fmt.Printf("\n The result is: %v", result)
fmt.Printf("\n The value of thing1 is: %v", thing1)

}


func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\n The memory location of the thing2 array is: %p\n", thing2)
	for i:= range thing2{
		thing2[i] = thing2[i]* thing2[i]
	}
	return *thing2
}