package main

import (
	"fmt"
	"log"
	"net"

	pb "hgndgn/grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterSortingService(server, &pb.SortingService{
		StringSort:       StringSort,
		Int32Sort:        Int32Sort,
		PersonSortByAge:  PersonSortByAge,
		PersonSortByName: PersonSortByName,
	})

	fmt.Printf("Listening on port %v\n", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
