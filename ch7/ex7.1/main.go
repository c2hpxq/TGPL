package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type WordCounter int

func (wc *WordCounter) Write(b []byte) (int, error){
	t := 0
	for i:= 0; i < len(b); {
		adv, word, err := bufio.ScanWords(b[i:], true)
		if err != nil {
			return 0, err
		}
		fmt.Println(string(word))
		t += 1
		i += adv
	}
	*wc += WordCounter(t)
	return t, nil
}

type LineCounter struct {
	n int
}

func (lc *LineCounter) Write(b []byte) (int, error) {
	t := 0
	for i := 0; i < len(b); {
		adv, _, err := bufio.ScanLines(b[i:], true)
		i += adv
		if err != nil {
			return i, err
		}
		t += 1
	}
	lc.n += t
	return len(b), nil
}

func (lc LineCounter) Count() int {
	return lc.n
}

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var wc WordCounter
	wc.Write(b)
	var lc LineCounter
	lc.Write(b)
	fmt.Printf("#word: %d, #line: %d", wc, lc.Count())
}
