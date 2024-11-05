package handin4

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "handin4/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	proto.UnimplementedElectionServer
}

var nextPort int
var client proto.ElectionClient

func (s *server) SendToken(ctx context.Context, token *proto.Token) (*proto.Empty, error) {
	fmt.Println("Do i have token? ", token.HasToken)
	fmt.Println("Doing task...")

	time.Sleep(2 * time.Second)

	fmt.Println("Done with task")

	if client != nil {
		fmt.Println("Sending token to client on port", nextPort)

		_, err := client.SendToken(ctx, token)
		if err != nil {
			log.Fatalf("Failed to send token: %v", err)
		}
	} else {
		findclient()

		fmt.Println("Sending token to client on port", nextPort)

		_, err := s.SendToken(ctx, token)
		if err != nil {
			return nil, err
		}
	}

	return &proto.Empty{}, nil
}

func StartClient(myPort int, _nextPort int) {
	nextPort = _nextPort

	fmt.Printf("Client running on port %d, next port is %d\n", myPort, nextPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", myPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterElectionServer(s, &server{})

	if myPort == 5050 {
		go findclient()
	}

	fmt.Printf("gRPC server listening on %v\n", myPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func findclient() {
	fmt.Println("Looking for client")

	var conn *grpc.ClientConn
	var err error

	for {
		conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", nextPort), grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			fmt.Println("Connected to client on port", nextPort)
			break
		}
	}

	client = proto.NewElectionClient(conn)

	_, err = client.SendToken(context.Background(), &proto.Token{HasToken: true})
	if err != nil {
		log.Fatalf("Failed to send token: %v", err)
	}

	fmt.Println("Token sent to client on port", nextPort)
}
