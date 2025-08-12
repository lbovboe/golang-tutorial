// the package name must be main
// so can go run main.go and go build can work
package main 

// stand for format
// formatting text for output or parsing text for input.
import "fmt" 


func main(){
	// ====== Printing =========
	fmt.Print("This is fmt.Print (no newline) | ")
	fmt.Println("This is fmt.Println (with newline)")
	fmt.Printf("This is fmt.Printf with formatting: %d + %d = %d\n",2,3,2+3)

	// ====== Creating a string with Sprintf ======
	// := assignment with auto infered type
	name := "Paul"
	message := fmt.Sprintf("hello, %s! Welcome to Go.", name)
	fmt.Println(message)

	// ====== Reading input =======
	// var is using when you want an explicit type
	// var is used when you want a null as the default value and assign the value later
	// normally don't use in function unless you not sure that is the initial value

	var firstName string
	var lastName string
	var age int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName) // reads until whitespace

	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName) // reads until newline

	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age) // reads formatted input and convert to number using %d

	// Output collected data
	fmt.Printf("Hi %s %s, you are %d years old!\n", firstName,lastName,age)


}
