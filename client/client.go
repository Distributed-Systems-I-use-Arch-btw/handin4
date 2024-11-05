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

func (s *server) SendToken(ctx context.Context, token *proto.Token) (*proto.Empty, error) {
	fmt.Println("Do i have token? ", token.HasToken)

	return &proto.Empty{}, nil
}

func StartClient(myPort int, nextPort int) {
	fmt.Printf("Client running on port %d, next port is %d\n", myPort, nextPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", myPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterElectionServer(s, &server{})

	go findclient(myPort, nextPort)

	fmt.Printf("gRPC server listening on %v\n", myPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func findclient(myPort int, nextPort int) {
	fmt.Println("Looking for client")

	var conn *grpc.ClientConn
	var err error

	for {
		conn, err = grpc.NewClient(fmt.Sprintf("localhost:%d", myPort))
		if err == nil {
			fmt.Println("Connected to client on port", nextPort)
			break
		}

		fmt.Printf("Failed to connect to client on port %d; retrying in 2 seconds...\n", nextPort)
		time.Sleep(2 * time.Second)
	}

	client := proto.NewElectionClient(conn)

	if myPort == 5050 {
		_, err = client.SendToken(context.Background(), &proto.Token{HasToken: true})
		if err != nil {
			log.Fatalf("Failed to send token: %v", err)
		}
		fmt.Println("Token sent to client on port", nextPort)
	}
}
