package main

import "fmt"

func main (){
	intArr := [...]int{1,2,3}
	fmt.Println(intArr) //print [1,2,3]
	intArr[1] =123
	fmt.Println(intArr) //print [1,123,3]
	
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice) //print [4,5,6]
	fmt.Printf("The length is %v and the capacity is  %v\n", len(intSlice), cap(intSlice))		 //"The length is 3 and the capacity is  3", 
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v and the capacity is  %v\n",  len(intSlice), cap(intSlice))         //The length is 4 and the capacity is  6", 
	fmt.Println(intSlice) //print [4,5,6,7]
	fmt.Println(intSlice[2]) // print ß
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...) // this append multiple values to the slice 
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3,8) // this will specify the length and capacity of the size here length is 3 and capacity is 8
	fmt.Println(intSlice3)
	fmt.Printf("The length is %v and the capacity is  %v\n", len(intSlice3), cap(intSlice3)) // The length is 3 and the capacity is  8


	//immediately initializing the array after declaring 
	// var newArr [3]int32 = [3]int32{1,2,3}  
	newArr := [...]int32{1,2,3} 
	//newArr := [3]int32{1,2,3}  // same as above
	fmt.Println(newArr)

	var myMap = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"]) //print 23
	fmt.Println(myMap2["Jason"]) //print 0 because it does not exits 

	//var age, ok = myMap2["Jason "]

	 // this delete the adam key from the array
	 //delete (myMap2, "Adam")

	// if ok {
	// 	fmt.Printf("the age is %v", age )
	// }else {
	// 	fmt.Printf("Invalid Name ")
	// }




	//To iterate with loops 
	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v \n" , name , age) 
	}

	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

// snipppet 1
	for i:=0; i<10;i++ {
		fmt.Println(i)
	}
	
// snippet 2	
	var i int = 0
	for {
		if i >= 10{
			break
		}
       fmt.Println(i)
       i = i + 1   
	}


}
 