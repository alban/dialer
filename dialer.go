package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	url := os.Args[1]
	numConn, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error with second argument\n")
		os.Exit(1)
	}

	fmt.Printf("Establishing %d TCP connections to %s\n", numConn, url)
	for x := 0; x < numConn; x++ {
		_, err := net.Dial("tcp", url)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	select {}
}
