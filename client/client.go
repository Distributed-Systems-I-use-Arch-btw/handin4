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
	"time"
)

var ports = []int{5050, 5051, 5052}

type server struct {
	proto.UnimplementedElectionServer
	ownPort  int
	hasToken bool
	nr       int
}

func (s *server) SendToken(ctx context.Context, ms *proto.Token) (*proto.Empty, error) {
	fmt.Println("got: ", ms.Boolean)

	if ms.Boolean {
		s.hasToken = true
		fmt.Println("DOING SOME COMPLICATED DATA MANIPULATION!")
		s.nr++
		time.Sleep(5 * time.Second)
	}

	return &proto.Empty{}, nil
}

func (s *server) client() {
	nextId := (s.ownPort + 1) % len(ports)
	err := s.startClient(nextId)
	for err != nil {
		err = s.startClient(nextId)
	}
	time.Sleep(2 * time.Second)
	s.client()
}

func (s *server) startClient(port int) error {
	conn, err := grpc.NewClient("localhost:"+strconv.Itoa(ports[port]), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := proto.NewElectionClient(conn)
	//TODO: BOOLEANS DOESN'T CORRECTLY!
	_, err = client.SendToken(context.Background(), &proto.Token{
		Boolean: s.hasToken,
		Data:    int32(s.nr),
	})
	if err != nil {
		return err
	}

	s.hasToken = false

	return nil
}

func (s *server) StartServer(port int) {
	gRPCserver := grpc.NewServer()

	log.Println("Server started")

	netListener, err := net.Listen("tcp", ":"+strconv.Itoa(ports[port]))
	if err != nil {
		panic(err)
	}

	proto.RegisterElectionServer(gRPCserver, s)

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
	if input == 0 {
		s.hasToken = true
	} else {
		s.hasToken = false
	}

	go s.client()
	go s.Server(input)

	fmt.Printf("Hello and welcome, %d!\n", input)

	for {

	}
}
