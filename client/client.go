package handin4

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "handin4/grpc"
	"log"
	"strconv"
)

var ports = []int32{5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057}

type Client struct {
	proto.UnimplementedElectionServer
}

func (c *Client) GetMessage(e *proto.Empty) (*proto.Message, error) {
	return &proto.Message{}, nil
}

func (c *Client) SendMessage(ms *proto.Message) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (c *Client) StartConnection(port *proto.Port) {
	nextClient, err := grpc.NewClient("localhost:"+strconv.Itoa(int(port.GetPort())), grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Hello, World!")
	fmt.Println(nextClient)
}

func StartClient(input int) {
	fmt.Printf("Hello and welcome, %d!\n", ports[input])

	if input != 0 {
		conn, err := grpc.NewClient("localhost:"+strconv.Itoa(int(ports[input+1])), grpc.WithInsecure())
		if err != nil {
			log.Fatalln(err)
		}

		nextClient := proto.NewElectionClient(conn)
		port := &proto.Port{Port: ports[input-1]}
		nextClient.StartConnection(context.Background(), port)
	}

	for {

	}
}
