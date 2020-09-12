package main

import (
	"context"
	"sort"
	"strings"

	pb "hgndgn/grpc/proto"
)

func sortString(str string) string {
	// alternative sort:
	// 		splitted := strings.Split(in.GetValue(), "")
	//		sort.Strings(splitted)
	//		sorted := strings.Join(splitted, "")

	var runes []string
	for _, v := range str {
		runes = append(runes, string(v))
	}
	sort.Strings(runes)
	return strings.Join(runes, "")
}

func StringSort(ctx context.Context, in *pb.StringInput) (*pb.StringResult, error) {
	result := sortString(in.GetValue())
	return &pb.StringResult{Value: result}, nil
}

func Int32Sort(ctx context.Context, in *pb.Int32Input) (*pb.Int32Result, error) {
	values := in.GetValues()
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return &pb.Int32Result{Values: values}, nil
}

func PersonSortByAge(ctx context.Context, in *pb.PersonList) (*pb.PersonList, error) {
	persons := in.GetPersons()
	sort.Slice(persons, func(i, j int) bool { return persons[i].Age < persons[j].Age })
	return &pb.PersonList{Persons: persons}, nil
}

func PersonSortByName(ctx context.Context, in *pb.PersonList) (*pb.PersonList, error) {
	persons := in.GetPersons()
	sortFunc := func(i, j int) bool { return strings.Compare(persons[i].Name, persons[j].Name) < 0 }
	sort.Slice(persons, sortFunc)
	return &pb.PersonList{Persons: persons}, nil
}
