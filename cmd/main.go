package main

import "fmt"

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
