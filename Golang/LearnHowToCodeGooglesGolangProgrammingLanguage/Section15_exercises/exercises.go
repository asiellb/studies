package main

import (
	"fmt"
)

type person struct { //The struct has to be defined outside the main function!
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello I am", p.first, p.last, "And I am", p.age, "years old, mentally!")
}

type square struct {
	side_len float64
	area     float64
}

func (s square) calculate_area() float64 {
	return s.side_len * s.side_len
}

type circle struct {
	diameter float64
	area     float64
}

func foo() int {
	return 42
}
func bar() (int, string) {
	return 42, "Meaning of everything"
}

func int_sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
func deferred_function() {
	fmt.Println("Dummy function to show the usage of the defer function")
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//Exercise 1
	//func, receiver, identifier, params, returns
	fmt.Println(foo())
	fmt.Println(bar())

	//Exercise 2
	fmt.Println("Sum of the passed numbers to int_sum:", int_sum(22, 33, 44))

	//Exercise 3
	//defer deferred_function() //Runs at the end of the program
	fmt.Println("One")
	deferred_function()
	fmt.Println("Two")
	fmt.Println("Three")

	//Exercise 4
	dani := person{
		first: "dani",
		last:  "verydani",
		age:   7,
	}
	fmt.Println(dani)
	dani.speak()

	//Exercise 5
	//meh, boring

	//Exercise 7: Assign func to variable, then call that func
	ai := intSeq()
	fmt.Println(ai())
	fmt.Println(ai())
	fmt.Println(ai())

}
