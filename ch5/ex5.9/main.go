package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	foo = "foo"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(expand(string(b), f))
}

func expand(s string, f func(string) string) string {
	ffoo := f(foo)
	return strings.Join(strings.Split(s, foo), ffoo)
}

func f(s string) string {
	return "goo"
}