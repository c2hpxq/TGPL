package main

import "fmt"

func varmax(args ...int) (max int) {
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return
}

func varmax1(a int, args ...int) (max int) {
	max = a
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return
}

func main() {
	fmt.Println(varmax())
	fmt.Println(varmax1(1, 2, 3))
	a := []int{20: 1}
	fmt.Println(a)
	fmt.Println(varmax(a...))
	fmt.Printf("%T", varmax1)
}