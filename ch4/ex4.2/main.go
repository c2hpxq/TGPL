package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var popcount [256]int8
var flagAlgo = flag.String("algo", "sha256", "choose between sha256/384/512")

func init() {
	for i := range popcount {
		popcount[i] = popcount[i/2] + int8(i&1)
	}
}

func main() {
	flag.Parse()
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch *flagAlgo {
	case "sha256":
		fmt.Printf("Result for %s = %x, Type = %[2]T\n", *flagAlgo, sha256.Sum256(content))
	case "sha384":
		fmt.Printf("Result for %s = %x, Type = %[2]T\n", *flagAlgo, sha512.Sum384(content))
	case "sha512":
		fmt.Printf("Result for %s = %x, Type = %[2]T\n", *flagAlgo, sha512.Sum512(content))
	default:
		fmt.Printf("No such algorithm %s, rerun with -algo sha256/sha384/sha512\n", *flagAlgo)
	}

}
