package main

import (
	"fmt"
	start "handin4/client"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println(len(args))
		fmt.Println("Too many arguments")
		os.Exit(1)
	}

	port, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Not an integer")
		os.Exit(1)
	}

	if port > 3 || port < 0 {
		fmt.Printf("Out of max range of %d\n", 3)
		os.Exit(1)
	}

	start.Start(port)
}
