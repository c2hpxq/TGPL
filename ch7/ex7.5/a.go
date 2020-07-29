package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type lReader struct {
	r io.Reader
	n int64
}

func (lr *lReader) Read(p []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if lr.n < int64(len(p)) {
		p = p[:lr.n]
	}
	n, err = lr.r.Read(p)
	lr.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &lReader{r: r, n: n}
}

func main() {
	lr := LimitReader(os.Stdin, 10)
	fmt.Println("start")
	p, err :=ioutil.ReadAll(lr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p))
}