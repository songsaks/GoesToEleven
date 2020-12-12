package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most and learning go programming language.")

	foo()

	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
