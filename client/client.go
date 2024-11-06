package handin4

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "handin4/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

type server struct {
	proto.UnimplementedElectionServer
	myPort      int
	nextPort    int
	client      proto.ElectionClient
	wantsAccess bool
	logger      *log.Logger
}

func (s *server) AccessCriticalSection() {
	fmt.Println("Doing task...")
	s.logger.Println("Doing task...")

	time.Sleep(2 * time.Second)

	fmt.Println("Done with task")
	s.logger.Println("Done with task")

	s.wantsAccess = false
}

func (s *server) SendToken(ctx context.Context, token *proto.Token) (*proto.Empty, error) {
	fmt.Printf("I, %d, have the token!\n", s.myPort)
	s.logger.Printf("I, %d, have the token!\n", s.myPort)

	if s.wantsAccess {
		s.AccessCriticalSection()
	}

	if s.client != nil {
		fmt.Println("Sending token to client on port", s.nextPort)
		s.logger.Println("Sending token to client on port", s.nextPort)

		_, err := s.client.SendToken(ctx, token)
		if err != nil {
			log.Fatalf("Failed to send token: %v", err)
		}
	} else {
		s.Findclient()

		fmt.Println("Sending token to client on port", s.nextPort)
		s.logger.Println("Sending token to client on port", s.nextPort)

		_, err := s.SendToken(ctx, token)
		if err != nil {
			return nil, err
		}
	}

	return &proto.Empty{}, nil
}

func StartClient(myPort int, nextPort int) {
	logFile, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger := log.New(logFile, "", log.LstdFlags)

	fmt.Printf("Client running on port %d, next port is %d\n", myPort, nextPort)
	logger.Printf("Client running on port %d, next port is %d\n", myPort, nextPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", myPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := &server{
		myPort:      myPort,
		nextPort:    nextPort,
		client:      nil,
		wantsAccess: false,
		logger:      logger,
	}

	g := grpc.NewServer()
	proto.RegisterElectionServer(g, s)

	if myPort == 5050 {
		go s.Findclient()
	}

	go s.DoIWantAccess()

	fmt.Printf("gRPC server listening on %v\n", myPort)
	logger.Printf("gRPC server listening on %v\n", myPort)
	if err := g.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Findclient() {
	fmt.Println("Looking for client")
	s.logger.Println("Looking for client")

	var conn *grpc.ClientConn
	var err error

	for {
		conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", s.nextPort), grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			fmt.Println("Connected to client on port", s.nextPort)
			s.logger.Println("Connected to client on port", s.nextPort)
			break
		}
	}

	s.client = proto.NewElectionClient(conn)

	_, err = s.client.SendToken(context.Background(), &proto.Token{})
	if err != nil {
		log.Fatalf("Failed to send token: %v", err)
	}

	fmt.Println("Token sent to client on port", s.nextPort)
	s.logger.Println("Token sent to client on port", s.nextPort)
}

func (s *server) DoIWantAccess() {
	if rand.Intn(5) == 1 && !s.wantsAccess {
		s.wantsAccess = true

		fmt.Printf("Port %d wants access\n", s.myPort)
		s.logger.Printf("Port %d wants access\n", s.myPort)
	}

	time.Sleep(2 * time.Second)

	s.DoIWantAccess()
}
