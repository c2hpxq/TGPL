package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var strs [100000]string
	for i := range strs {
		strs[i] = "aaaaa"
	}

	for n := 100; n <= len(strs); n *= 10 {
		start := time.Now()
		res := strings.Join(strs[:n], "")
		t := time.Since(start).Milliseconds()
		fmt.Println("join, form string of length ", len(res), ", cost time ", t, " ms")

		start = time.Now()
		res = ""
		for _, arg := range strs[:n] {
			res += arg
		}
		t = time.Since(start).Milliseconds()
		fmt.Println("concat 1 by 1, form string of length ", len(res), ", cost time ", t, " ms")

	}

}