package handin4

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "handin4/grpc"
	"log"
	"net"
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

	fmt.Printf("gRPC server listening on %v\n", myPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
