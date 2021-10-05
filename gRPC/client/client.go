package main

import (
	"context"
	"fmt"
	pb "itucm/grpc"

	"google.golang.org/grpc"
)

func main() {
	// gRPC channel
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		panic("Could not connect!")
	}
	defer conn.Close()

	// client stub
	client := pb.NewITUCourseManagerClient(conn)

	// client implementation
	courses, err := client.GetAllCourses(context.Background(), &pb.Empty{})
	if err != nil {
		panic("Could not get all courses:" + err.Error())
	}

	fmt.Println(courses.Courses)
}