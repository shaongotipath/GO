package main

import "fmt"

type Person struct {
	name   string
	age    int
	salary float32
	isMale bool
}

func increment(n *int) {
	*n = *n + 1
}

func worker(done chan bool) {
	fmt.Println("Working...")
	done <- true
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	value := <-ch

	fmt.Println("Received:", value)

	done := make(chan bool)

    fmt.Println("Work is not done!")

	go worker(done)

	<-done
    
	fmt.Println("Work is done!")

	n := 5
	fmt.Println("Before increment:", n)

	increment(&n)

	fmt.Println("After increment:", n)

	var a = map[string]int{"John": 25, "Jane": 30, "Sam": 35, "Tom": 40, "Alice": 45}

	fmt.Println(a)
	for key, value := range a {
		fmt.Printf("Name: %s, Age: %d\n", key, value)
	}

	company := 10
	for i := 0; i < company; i++ {
		if i == 5 {
			fmt.Println("company ", 5, " is founded")
		}
		switch i {
		case 1:
			fmt.Println("Company 1 is in startup phase")
		case 2:
			fmt.Println("Company 2 is growing")
		case 3:
			fmt.Println("Company 3 is established")
		case 4:
			fmt.Println("Company 4 is expanding")
		case 5:
			fmt.Println("Company 5 is a market leader")
		default:
			fmt.Println("Company", i, "is in an unknown phase")
		}
		fmt.Println("company", i)
	}

	var users = [...]string{"John", "Jane", "Sam", "Tom", "Alice"}

	users2 := [...]string{"John", "Jane", "Sam", "Tom", "Alice"}

	fmt.Println(len(users))
	fmt.Println(users2)

	person := Person{
		name:   "John",
		age:    25,
		salary: 5000.00,
		isMale: true,
	}

	person2 := Person{
		name:   "Jane",
		age:    30,
		salary: 6000.00,
		isMale: false,
	}

	fmt.Println("1st Person Information--------")
	fmt.Println("Name : ", person.name)
	fmt.Println("Age : ", person.age)
	fmt.Println("Salary : ", person.salary)
	fmt.Println("Male : ", person.isMale)

	fmt.Println("2nd Person Information--------")
	fmt.Println("Name : ", person2.name)
	fmt.Println("Age : ", person2.age)
	fmt.Println("Salary : ", person2.salary)
	fmt.Println("Male : ", person2.isMale)

	x := 10

	var ptr *int

	ptr = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of ptr:", *ptr)
	fmt.Println("Address stored in ptr:", ptr)

}
