package handin4

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "handin4/grpc"
	"log"
	"net"
	"strconv"
)

var ports = []int{5050, 5051, 5052, 5053, 5054, 5055, 5056, 5057}

type server struct {
	proto.UnsafeByzantiumServer
}

func (c *server) GetMessage(ctx context.Context, e *proto.Empty) (*proto.Message, error) {
	return &proto.Message{}, nil
}

func (c *server) SendMessage(ctx context.Context, ms *proto.Message) (*proto.Empty, error) {
	fmt.Println(ms.Message)
	return &proto.Empty{}, nil
}

func client() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewByzantiumClient(conn)

	_, _ = client.SendMessage(context.Background(), &proto.Message{Message: "PEtAR"})
}

func (s *server) StartServer(port int) {
	gRPCserver := grpc.NewServer()

	log.Println("Server started")

	netListener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}

	proto.RegisterByzantiumServer(gRPCserver, s)

	err = gRPCserver.Serve(netListener)
	if err != nil {
		panic(err)
	}

}

func (s *server) Server(port int) {
	s.StartServer(port)
}

func StartClient(input int) {
	s := &server{}
	s.Server(input)

	if input == 5051 {
		client()
	}

	fmt.Printf("Hello and welcome, %d!\n", input)
}
