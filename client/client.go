package handin4

import (
	"fmt"
	proto "handin4/grpc"
)

var ports = []int{5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057}

type Client struct {
	proto.UnsafeByzantiumServer
}

func (c *Client) GetMessage(e *proto.Empty) (*proto.Message, error) {
	return &proto.Message{}, nil
}

func (c *Client) SendMessage(ms *proto.Message) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func StartClient(input int) {

	fmt.Printf("Hello and welcome, %d!\n", input)
}
