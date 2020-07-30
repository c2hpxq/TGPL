package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		tmp := strings.Split(arg, "=")
		loc, addr := tmp[0], tmp[1]
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		go mustCopy(loc, os.Stdout, conn)
	}

	for {

	}
}

func mustCopy(loc string, dst io.Writer, src io.ReadCloser) {
	defer src.Close()
	reader := bufio.NewReader(src)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(loc, ":", s)
	}

}