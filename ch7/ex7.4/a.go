package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type sReader struct {
	s string
	i int64
}

func (sr *sReader) Read(p []byte) (n int, err error) {
	if sr.i >= int64(len(sr.s)) {
		return 0, io.EOF
	}
	n = copy(p, sr.s[sr.i:])
	sr.i += int64(n)
	return
}

func newReader(s string) *sReader {
	return &sReader{s: s, i: 0}
}

func main() {
	p, err := ioutil.ReadAll(newReader("abc"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(p))
}
