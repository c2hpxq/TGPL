package main

import (
	"fmt"
	"strings"
)

func join(sep string, ss ...string) string {
	return strings.Join(ss, sep)
}

func main() {
	fmt.Println(join(",", []string{"aa", "bb"}...))
}