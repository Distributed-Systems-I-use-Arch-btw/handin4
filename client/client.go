package handin4

import (
	"fmt"
	"google.golang.org/grpc"
	proto "handin4/grpc"
	"log"
	"strconv"
)

var ports = []int{5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057}

type Client struct {
	proto.UnimplementedElectionServer
}

func (c *Client) GetMessage(e *proto.Empty) (*proto.Message, error) {
	return &proto.Message{}, nil
}

func (c *Client) SendMessage(ms *proto.Message) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (c *Client) StartConnection(port *proto.Port) (*proto.Empty error) {
	nextClient, err := grpc.NewClient("localhost:" + port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	return &proto.Empty{}, nil
}

func StartClient(input int) {
	fmt.Printf("Hello and welcome, %d!\n", ports[input])

	if input != 0 {
		conn, err := grpc.NewClient("localhost:"+strconv.Itoa(ports[input]), grpc.WithInsecure())
		if err != nil {
			log.Fatalln(err)
		}

		nextClient := proto.NewElectionClient(conn)
		nextClient.StartConnection(ports[input])
	}
}