package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	url := os.Args[1]

	fmt.Printf("Connecting to %s\n", url)
	for x := 0; x < 10; x++ {
		_, err := net.Dial("tcp", url)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	select {}
}
