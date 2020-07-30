// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/google/syzkaller/pkg/osutil"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleConn(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(cmd)
		cmdArgs := strings.Split(cmd[:len(cmd)-1], " ")
		switch cmdArgs[0] {
		case "ls":
			filenames, err := osutil.ListDir(".")
			if err != nil {
				continue
			}
			fmt.Fprintf(c, strings.Join(filenames, "\n"))
		case "get":
			f, err := os.Open(cmdArgs[1])
			if err != nil {
				continue
			}
			io.Copy(c, f)
		case "close":
			return
		}
	}
}

var port = flag.String("port", "8000", "listening port")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ftp starts serving.")
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
