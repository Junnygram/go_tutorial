# Go essentials

`func functioName(param1 type1, param2 type2) return Type{
//Function body
return result
}`

Concurrency is one of the defining features of Go. Go provides native support for concurrency through goroutines and channels. Goroutines are lightweight threads that allow developers to execute concurrent tasks efficiently. Channels facilitate communication and synchronisation between goroutines, making it easier to build concurrent systems.

Adding visibility to identifiers
Identifiers include variable names, custom data types, functions, and other constructs in Go. Go uses capitalisation to determine if a package-level identifier is visible outside of the package where it is declared.
An identifier whose Name starts with an uppercase letter is exported. An identifier whose Name begins with a lowercase letter or underscore can only be accessed from within the package where it is declared.
Note that Go developers use camel case for naming identifiers.

`var Name string // can be accessed anywhere in the codebase.
var localName string // can only be accessed in the source code it is declared in. `

The go Command

The Go command is a fundamental tool in Go that serves various purposes, including compiling and running Go programs, managing packages, and performing various development tasks. Here are some common go commands:

`var Name string // can be accessed anywhere in the codebase.
var localName string // can only be accessed in the source file it is declared in. `

go build : Compiles the Go program and generates an executable binary. If there are no errors, this will create an executable le with the same name as the directory.
go run : Compiles and runs a Go program in a single step. This is useful for quickly testing small programs.
go install : Compiles and installs the Go package or program. If the package is a command (an executable), it will be placed in the bin directory dened in your GOPATH.
go get: : Fetches and installs packages from the remote repository.
go test : Runs tests associated with the current package.

Variable and Constant Declaration
To declare a variable, use the var keyword followed by the variable name, its type, and an optional initial value. If you don't specify the value, the Go compiler will initialize it to the zero value of that variable type.

`package main
import "fmt"
func main() {
var name string // initialized as an empty string:""  
name = "John"// name string now becomes "John" fmt.println("Name:", name) // Output: Name: John
var ageint=30fmt.Println("Age:", age) // Output: Age: 30// You can also use the short variable declaration syntax City := "New York"fmt.Println("City:", city) // Output: City: New York
} `

In a backend system, you often need to work with different data types. You declare variables explicitly, indicating their types for clarity and type safety.

`package main
import "fmt"
func main() {var databaseURL stringdatabaseURL = "mongodb://localhost:27017/mydb" fmt.Println("Database URL:", databaseURL)
var port int = 8080
fmt.Println("Port:", port)
// Using the short variable declaration syntax
serverName := "Backend Server"
fmt.Println("Server Name:", serverName)
}

`

Multiple Variable Declaration
Go allows you to declare multiple variables in a single statement. You can do variable declaration in two ways and in both local and global scopes. You can also do a short assignment declaration with multiple variables.

`package main 
import "fmt" 
func main() { 
var( 
maxConnections int = 100 
timeoutSeconds int = 30 ) 
fmt.Println("Max Connections:", maxConnections) 
fmt.Println("Timeout Seconds:", timeoutSeconds) 
}`

Constant Declaration
To declare a constant, use the const keyword followed by the constant Name, its type (optional), and the constant's value. Constants are preferably declared globally (outside of functions) but can be declared globally. Unlike variables, constants must be initialized at the time of declaration and cannot be changed later.

`package main 
import "fmt" 
func main() { 
const ( 
MaxRetries = 3 
BuerSize =1024 
RequestType = "POST" 
) 
fmt.Println("Max Retries:", MaxRetries) 
fmt.Println("Buer Size:", BuerSize) 
fmt.Println("Request Type:", RequestType) 
}`

Enumerated Constants
Go also supports enumerated constants using the iota keyword. Enumerated constants are auto-incremented integer values commonly used to dene a set of related named constants. For example, say we have a backend system that allows users to subscribe to services; here is what our subscription model will look like:

`import "fmt"
func main() { const (
Basic = iota // 0
)
Premium Gold Enterprise
// 1 (automatically incremented from iota) // 2 (automatically incremented from iota)
// 3 (automatically incremented from iota)
fmt.Println("Available subscription packages and index:") fmt.Println("Basic:", Basic)fmt.Println("Premium:", Premium)fmt.Println("Gold:", Gold)
fmt.Println("Enterprise:", Enterprise) }`
