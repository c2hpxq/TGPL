package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	writer io.Writer
	counter *int64
}

func (cw countingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	*cw.counter += int64(n)
	return
}

func CountingWriter(w io.Writer) (writer io.Writer, counter *int64) {
	var i int64
	cw := countingWriter{writer: w, counter: &i}
	return cw, &i
}

func main() {
	writer, counter := CountingWriter(os.Stdout)
	writer.Write([]byte("aaaaa"))
	fmt.Println(*counter)
	fmt.Println(*counter)
}