package main

import (
	"fmt"
	proto "handin4/grpc"
)

type client struct {
	proto.UnimplementedChittyChatServer
}

func main() {

	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {

		fmt.Println("i =", 100/i)
	}
}
