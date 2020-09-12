package main

import (
	"context"
	"fmt"
	"time"

	pb "hgndgn/grpc/proto"

	"google.golang.org/grpc"
)

const (
	host    = "localhost"
	port    = ":8000"
	timeout = 10
)

func main() {
	conn, _ := grpc.Dial(
		host+port,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(timeout*time.Second),
	)
	defer conn.Close()

	c := pb.NewSortingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	const input = "zdefaxv"
	sortedInput, _ := c.StringSort(ctx, &pb.StringInput{Value: input})

	ints := []int32{75, 3, 476, 2, 26, 56, 0}
	sortedInts, _ := c.Int32Sort(ctx, &pb.Int32Input{Values: ints})

	persons := &pb.PersonList{Persons: []*pb.Person{
		&pb.Person{Age: 23, Name: "John"},
		&pb.Person{Age: 21, Name: "Lisa"},
	}}
	sortedByAge, _ := c.PersonSortByAge(ctx, persons)
	sortedByName, _ := c.PersonSortByName(ctx, persons)

	fmt.Printf("Sorted string:  \t%v\n", sortedInput.GetValue())
	fmt.Printf("Sorted integers: \t%v\n", sortedInts.GetValues())
	fmt.Printf("Sorted by 'age': \t%v\n", sortedByAge.GetPersons())
	fmt.Printf("Sorted by 'name':\t%v\n", sortedByName.GetPersons())
}
