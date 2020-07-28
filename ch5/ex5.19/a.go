package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (res int) {
	type retVal struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case retVal{}:
			res = 1
		default:
			panic(p)
		}
	}()
	panic(retVal{})
}