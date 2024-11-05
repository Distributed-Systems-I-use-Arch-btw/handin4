package main

import (
	"fmt"
	start "handin4/client"
	"os"
	"strconv"
)

var ports = []int{5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057}

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println(len(args))
		fmt.Println("Too many arguments")
		os.Exit(1)
	}
	nr, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not an integer")
		os.Exit(1)
	}
	if nr > len(ports) || nr < 1 {
		fmt.Printf("Out of max range of &d\n", len(ports))
		os.Exit(1)
	}
	start.StartClient(ports[nr])
}
