package main

import (
	"context"
	"fmt"
	pb "itucm/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Course struct {
	Id           int32     `json:"id"`
	Name         string  `json:"name"`
	Workload     int32     `json:"workload"`
	Satisfaction float32 `json:"satisfaction"`
	Teacher      int32     `json:"teacher"`
	Students     []int32   `json:"students"`
}

var initialCourses = []Course{
	{
		Id:           0,
		Name:         "Basic basics",
		Workload:     15,
		Satisfaction: 0.75,
		Teacher:      0,
		Students:     []int32{0, 1, 2},
	},
	{
		Id:           1,
		Name:         "Rudimentary roots 101",
		Workload:     7,
		Satisfaction: 0.95,
		Teacher:      1,
		Students:     []int32{2, 3, 4},
	},
}


type ITUCourseManagerServer struct {
	pb.UnimplementedITUCourseManagerServer
	courses []pb.Course
}

// Implement gRPC functions

func (s *ITUCourseManagerServer) GetAllCourses(ctx context.Context, empty *pb.Empty) (*pb.CourseList, error) {
	hardcodedoutput := pb.CourseList{
		Courses: []*pb.Course{
			&s.courses[0],
			&s.courses[1],
		},
	}
	return &hardcodedoutput, nil
}

func (s *ITUCourseManagerServer) PostCourse(ctx context.Context, course *pb.Course) (*pb.Empty, error) {
	panic("Not implemented!")
}

func (s *ITUCourseManagerServer) GetCourse(ctx context.Context, id *pb.Id) (*pb.Course, error) {
	panic("Not implemented!")
}

func (s *ITUCourseManagerServer) PutCourse(ctx context.Context, course *pb.Course) (*pb.Empty, error) {
	panic("Not implemented!")
}

func (s *ITUCourseManagerServer) DeleteCourse(ctx context.Context, id *pb.Id) (*pb.Empty, error) {
	panic("Not implemented!")
}


func newServer() *ITUCourseManagerServer {
	s := &ITUCourseManagerServer{
		courses: []pb.Course{
			{
				Id:           &initialCourses[0].Id,
				Name:         &initialCourses[0].Name,
				Workload:     &initialCourses[0].Workload,
				Satisfaction: &initialCourses[0].Satisfaction,
				Teacher:      &initialCourses[0].Teacher,
				Students:     initialCourses[0].Students,
			},
			{
				Id:           &initialCourses[1].Id,
				Name:         &initialCourses[1].Name,
				Workload:     &initialCourses[1].Workload,
				Satisfaction: &initialCourses[1].Satisfaction,
				Teacher:      &initialCourses[1].Teacher,
				Students:     initialCourses[1].Students,
			},
		},
	}
	return s
}

func main() {
	fmt.Println("STARTING SERVER")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Falied to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterITUCourseManagerServer(grpcServer, newServer())
	fmt.Println("SERVER STARTED")
	grpcServer.Serve(lis)
}