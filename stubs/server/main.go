package main

import (
	"context"
	"errors"
	"log"
	"net"
	"reflect"

	"github.com/ali2210/fictional-engine/stubs/grpcc"
	"google.golang.org/grpc"
)

type Plan int64

const (
	Deadline Plan = iota
	Busy
	Started
)

type Task_features struct {
	TaskName string
	Deadline string
	Level    Plan
	Id       int64
}

type taskRouteServer struct {
	grpcc.UnimplementedTaskPlan
}

var feature *Task_features

func StoreTask(task *Task_features) {
	feature = task
}

func GetTask() *Task_features { return feature }

func NewServer() *taskRouteServer {
	return &taskRouteServer{}
}

func (r *taskRouteServer) Add(ctx context.Context, t *grpcc.Create_Task) (*grpcc.Status, error) {

	if reflect.DeepEqual(t, grpcc.Create_Task{}) {
		return &grpcc.Status{Result: grpcc.Result_Error}, errors.New("empty task ")
	}
	return &grpcc.Status{Result: grpcc.Result_Ok}, nil
}

func (r *taskRouteServer) Remove(ctx context.Context, t *grpcc.Delete_Task) (*grpcc.Status, error) {

	if reflect.DeepEqual(t, grpcc.Delete_Task{}) {
		return &grpcc.Status{Result: grpcc.Result_Error}, errors.New("task object should not be empty")
	}
	return &grpcc.Status{Result: grpcc.Result_Ok}, nil
}

func (r *taskRouteServer) List(ctx context.Context, t *grpcc.View_Task) (*grpcc.Status, error) {

	if reflect.DeepEqual(t, grpcc.View_Task{}) {
		return &grpcc.Status{Result: grpcc.Result_Error}, errors.New("please create a task; ")
	}
	return &grpcc.Status{Result: grpcc.Result_Ok}, nil
}

var port string = ":9009"

func main() {

	log.Println("Starting server... Listening @:", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Failed to Listen:", err.Error())
		return
	}

	log.Println("GRPC Server started...")

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	log.Println("Service Request forwarding..... :")
	grpcc.RegisterService(server, NewServer())

	log.Println("Grpc server listening @ 9009......")
	server.Serve(listener)
}
