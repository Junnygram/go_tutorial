a module is a single file containing code, whereas a package is a collection of modules

# to compile and run in a single command

go run directory

# to print

import "fmt"

# dataTypes

bool
float32, float64
int, int16, int32, int64
rune
string
uint, uint8, uint16, uint32, uint64

# same code

`var myVar = "text"`
`myVar := "text"`

# to initialise two or more variables

`var var1, var2 int = 1, 2`
meaning var1= 1, var2 = 2 and can also be wriiten as

`var1 var2:=1, 2`

# if statement and switch statement available

# ARRAY

NB: the elemt in an array starts with the index 0

`var intArr [3]int32`
meaning this array holds 3 elements and cannot be changed after being initialized and also spicy that all the elements in the array are type int 32

`fmt.Println(intArr[1:3])`
this will access the element 1 and 2 in the array

`fmt.Println(&intArr[2])`
this prints out the memory location of index 2 and its store in a contiguous memory location

# all three line of codes are the same

`var newArr [3]int32 = [3]int32{1,2,3}  
	newArr := [...]int32{1,2,3} 
	newArr := [3]int32{1,2,3} `
0

#append multiple values to the slice by using spread operator
