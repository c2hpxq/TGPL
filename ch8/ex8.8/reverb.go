// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	fmt.Println("here")
	input := bufio.NewScanner(c)
	var w sync.WaitGroup
	defer func() {
		w.Wait()
		c.Close()
	}()

	lines := make(chan string)
	go func() {
		for input.Scan(){
			lines <- input.Text()
		}
	}()

	timer := time.NewTimer(10*time.Second)
	for {
		select {
		case line := <-lines:
			w.Add(1)
			go func() {
				echo(c, line, 1*time.Second)
				w.Done()
			}()
		case <-timer.C:
			fmt.Println("hanged for too long, break connection")
			return
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
