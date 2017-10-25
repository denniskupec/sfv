package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ieee := crc32.NewIEEE()

	if _, err := io.Copy(ieee, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%08X\n", ieee.Sum(nil))
}
