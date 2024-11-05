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

var ports = []int{5050, 5051, 5052}

type server struct {
	proto.UnsafeByzantiumServer
	ownPort int
}

func (c *server) GetMessage(ctx context.Context, e *proto.Empty) (*proto.Message, error) {
	return &proto.Message{}, nil
}

func (c *server) SendMessage(ctx context.Context, ms *proto.Message) (*proto.Empty, error) {
	fmt.Println(ms.Message)
	return &proto.Empty{}, nil
}

func (s *server) client() {
	err := startClient((s.ownPort + 1) % len(ports))
	for err != nil {
		err = startClient((s.ownPort + 1) % len(ports))
	}
}

func startClient(port int) error {
	conn, err := grpc.NewClient("localhost:"+strconv.Itoa(ports[port]), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := proto.NewByzantiumClient(conn)

	_, err = client.SendMessage(context.Background(), &proto.Message{Message: "PEtAR"})
	if err != nil {
		return err
	}

	return nil
}

func (s *server) StartServer(port int) {
	gRPCserver := grpc.NewServer()

	log.Println("Server started")

	netListener, err := net.Listen("tcp", ":"+strconv.Itoa(ports[port]))
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

func Start(input int) {
	s := &server{ownPort: input}

	go s.client()
	go s.Server(input)

	fmt.Printf("Hello and welcome, %d!\n", input)

	for {

	}
}
