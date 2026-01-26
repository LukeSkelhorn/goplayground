package main

import (
	"fmt"
	"math/rand"
)

func main() {
	types()
	forLoops()
}

func types() {
	// Basic types
	var num int = 7
	var str string = "Hello, World!"
	var boolVal bool = true
	var floatNum float64 = 3.14

	fmt.Println(num, str, boolVal, floatNum)
}

func forLoops() {
	fmt.Println("For Loops")

	fmt.Println("Compete for leep")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("variable decliration before loop")
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Leave off the increments if you have more complicated increment rule inside the loop")
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 { // Returns the division remainder
			i++
		} else {
			i += 2
		}
	}

	// outer loop example
	samples := []string{"hello", "workd", "from", "go"}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			// Returns the index, rung (Unicode code point), and string representation of each character in the string
			if r == ';' {
				continue outer
			}
		}
		fmt.Println()
	}
}

// shadowing example
// variable shadowing occurs when a variable declared within a certain scope
// (e.g., inside a function or block) has the same name as a variable declared
// in an outer scope. The inner variable "shadows" the outer variable within its scope.
func shadowTypes() {
	x := 7
	fmt.Println(x)
	if true {
		x := 17
		fmt.Println(x)
	}
	fmt.Println(x)
}

func switchSatements() {
	// switch statement example
	age := 28
	switch {
	case age < 13:
		fmt.Println("Child")
	case age >= 13 && age < 20:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}

	switch age {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		fmt.Println("Child")
	case 13, 14, 15, 16, 17, 18, 19:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}

	// for loop with switch and break example
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 5:
			fmt.Println("halfway there")
		case 7:
			fmt.Println("lucky number seven, breaking out of loop")
			break loop
		default:
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("fuzz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

func gotoStatement() {
	fmt.Println("goto statement example")
	// Create random number
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loops completes normally")
done:
	fmt.Println("do soemthing complex no matter why we left the loop")
	fmt.Println(a)
}
