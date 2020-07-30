// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				// in case client not ready even if buffer is already full, drop the msg
				select {
				case cli <- msg:
				default:
				}
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	// ex 8.15, broadcaster may block on single send if corresponding client doesn't read in time
	// use buffered channel here
	ch := make(chan string, 10) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	who := conn.RemoteAddr().String()
	// ex8.14, add user naming support, use ip:port if empty
	ch <- "Please enter your name(we use ip:port if empty): "
	input.Scan()
	if input.Text() != "" {
		who = input.Text()
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	// ex8.13 add timeout
	const timeout = 10*time.Second
	timer := time.NewTimer(timeout)
	userBreak := make(chan struct{})
	go func() {
		for input.Scan() {
			timer.Reset(timeout)
			messages <- who + ": " + input.Text()
		}
		close(userBreak)
	}()

	// tell from timeout & voluntary leave
	select {
	case <-timer.C:
		fmt.Println(who + " timeout")
	case <-userBreak:
		fmt.Println(who + " left")
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
