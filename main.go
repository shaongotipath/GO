// // Package main implements a simple Go program that demonstrates various basic functionalities
// // such as printing to the console, looping, function calls, array and slice initialization,
// // and basic arithmetic operations.
// package main

// import (
// 	"fmt"
// 	"math"
// 	// "net/http"
// )

// // // promoted prints a message indicating that the given person has been promoted.
// // func promoted(name string) {
// // 	fmt.Println(name, "just got promoted...!")
// // }

// // // salary prints a fixed salary amount.
// // func salary() {
// // 	fmt.Println("Salary is 1000")
// // }

// // // areaCalculation calculates and prints the area of a rectangle given its length and width.
// // func areaCalculation(length int, width int) int {
// // 	area := length * width
// // 	return area
// // }

// // func max(num1 int, num2 int) int {
// // 	var result int

// // 	if num1 > num2 {
// // 		result = num1
// // 	} else {
// // 		result = num2
// // 	}

// // 	return result
// // }

// // func swap(str1 string, str2 string) (string, string) {
// // 	return str2, str1
// // }

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Circle struct {
// 	radius float64
// }

// type Rectangle struct {
// 	length, width float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c Circle) Perimeter() float64 {
// 	return math.Pi * c.radius
// }

// func (r Rectangle) Area() float64 {
// 	return r.length * r.width
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }

// // Function to display shape details (polymorphism in action)
// func displayShapeInfo(s Shape) {
//     fmt.Printf("Shape Info:\n")
//     fmt.Printf("Area: %.2f\n", s.Area())
//     fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
// }

// // type Address struct {
// //     City    string
// //     ZipCode int
// // }

// // type Employee struct {
// //     Name    string
// //     Address
// // }

// // Handler function for the root path
// // func handler(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "Welcome to the Root Path!")
// // }

// // func aboutHandler(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "This is the About Page!")
// // }

// // func greetHandler(w http.ResponseWriter, r *http.Request) {
// // 	// Parse query parameter "name"
// // 	name := r.URL.Query().Get("name")
// // 	if name == "" {
// // 		name = "Stranger" // Default to "Stranger" if no name is provided
// // 	}

// // 	// Respond with a greeting
// // 	fmt.Fprintf(w, "Hello, %s!", name)
// // }

// func main() {
// 	// Register the handler for the root path
// 	// http.HandleFunc("/", handler)

// 	// http.HandleFunc("/about", aboutHandler)

// 	// http.HandleFunc("/greet", greetHandler)

// 	// // Start the server on port 8080
// 	// fmt.Println("Starting server on :8080...")
// 	// err := http.ListenAndServe(":8080", nil)
// 	// if err != nil {
// 	// 	fmt.Println("Error starting server:", err)
// 	// }

// 	// Create instances of Circle and Rectangle
// 	c := Circle{radius: 5}
// 	r := Rectangle{length: 10, width: 4}

// 	// Display details using the Shape interface
// 	fmt.Println("Circle:")
// 	displayShapeInfo(c)

// 	fmt.Println("\nRectangle:")
// 	displayShapeInfo(r)

// 	// fmt.Println("-------------------------")

// 	// // // Create an Employee instance
// 	// emp :=Employee{
// 	// 	Name: "John Doe",
// 	// 	Address: Address{City: "New York", ZipCode: 10001},
// 	// }
// 	// fmt.Println(emp);
// 	// fmt.Println(emp.Name)      // Output: Alice
// 	// fmt.Println(emp.City)      // Output: New York (Access embedded struct fields directly)
// 	// fmt.Println(emp.ZipCode)   // Output: 10001

// 	// // Prints "Hello, World!..." to the console.
// 	// fmt.Println("Hello, World!...")

// 	// // Loops from 0 to 10 in steps of 10 and prints the current value of i.
// 	// for i := 0; i <= 10; i += 10 {
// 	// 	fmt.Println(i)
// 	// }

// 	// // Calls the promoted function with the name "John".
// 	// promoted("John")

// 	// // Calls the salary function.
// 	// salary()

// 	// // Array initialization examples.
// 	// arr1 := [5]int{}              // Not initialized
// 	// arr2 := [5]int{1, 2}          // Partially initialized
// 	// arr3 := [5]int{1, 2, 3, 4, 5} // Fully initialized

// 	// arrL := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
// 	// arrH := [...]int{1, 2, 3, 4, 5, 6}

// 	// // Prints the length of arrL and arrH.
// 	// fmt.Println(len(arrL))
// 	// fmt.Println(len(arrH))

// 	// // Prints the contents of arr1, arr2, and arr3.
// 	// fmt.Println(arr1)
// 	// fmt.Println(arr2)
// 	// fmt.Println(arr3)

// 	// // Slice initialization examples.
// 	// myslice1 := []int{}
// 	// fmt.Println(len(myslice1)) // Prints the length of myslice1.
// 	// fmt.Println(cap(myslice1)) // Prints the capacity of myslice1.
// 	// fmt.Println(myslice1)      // Prints the contents of myslice1.

// 	// myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
// 	// fmt.Println(len(myslice2)) // Prints the length of myslice2.
// 	// fmt.Println(cap(myslice2)) // Prints the capacity of myslice2.
// 	// fmt.Println(myslice2)      // Prints the contents of myslice2.

// 	// // Variable x of type float32.
// 	// var x float32
// 	// x = 20.0
// 	// fmt.Println(x)                     // Prints the value of x.
// 	// fmt.Printf("x is of type %T\n", x) // Prints the type of x.

// 	// //find max number
// 	// var maxNumber int

// 	// maxNumber = max(10, 15)

// 	// fmt.Printf("Max value is : %d\n", maxNumber)

// 	// //swap two strings
// 	// str1, str2 := swap("Alan", "Bob")

// 	// fmt.Printf("str1 is %s, str2 is %s\n", str1, str2)

// 	// // Area calculation example.
// 	// var length, width int

// 	// fmt.Print("Enter the length of the rectangle: ")
// 	// fmt.Scan(&length)

// 	// fmt.Print("Enter the width of the rectangle: ")
// 	// fmt.Scan(&width)

// 	// area :=areaCalculation(length, width);

// 	// fmt.Printf("area of the rectangle is: %d\n", area)

// }
