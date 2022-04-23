package client

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/ali2210/fictional-engine/graph"
	"github.com/ali2210/fictional-engine/stubs/grpcc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Plan int64
type route int64

const (
	Create route = iota
	Delete
	View
)

type Task_features struct {
	TaskName string
	Deadline string
	Level    Plan
	Id       int64
}

var feature *Task_features

func _init(port string) (error, *grpc.ClientConn) {

	if port == "" {
		return errors.New("port is not specified :("), &grpc.ClientConn{}
	}

	log.Println("Client connect to server .....", port)
	connect, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.New("Dial error: " + err.Error()), connect
	}
	defer connect.Close()

	// if strings.Contains(t.TaskName, " ") {
	// 	return &grpcc.Status{Result: grpcc.Result_Error}, errors.New("task name is not specified")
	// }

	return err, connect
}

func Client_Init() error {

	err, client := _init("127.0.0.1:9009")
	if err != nil {
		log.Fatalln("Failed to Initialize Client: .... :( ", err.Error())
		return err
	}

	if reflect.DeepEqual(GetTask(), &Task_features{}) {
		log.Println("task is not initialized .... :(")
		return errors.New("task is not initialized .... :(")
	}

	if reflect.DeepEqual(GetTask().TaskName, " ") {
		log.Println("task name is not specified .... :(")
		return errors.New("task name is not specified .... :(")
	}

	task_created := &grpcc.Create_Task{TaskName: GetTask().TaskName, Deadline: GetTask().Deadline, Id: GetTask().Id}
	if GetTask().Level == 0 {
		task_created.LevelWork = grpcc.Priority_deadline
	}
	if GetTask().Level == 1 {
		task_created.LevelWork = grpcc.Priority_busy
	}
	if GetTask().Level == 2 {
		task_created.LevelWork = grpcc.Priority_started
	}

	log.Println("Cargo move forward  ..... :) : ")

	RouteServices(task_created, client)

	return nil
}

func StoreTask(task *Task_features) {
	feature = task
}

func GetTask() *Task_features { return feature }

func RouteServices(t *grpcc.Create_Task, c *grpc.ClientConn) {

	switch Webroute {
	case Create:
		err := add_route(t, c)
		if err != nil {
			log.Fatalln("Failed :", err.Error())
			return
		}
	case Delete:
		err := delete_route(t, c)
		if err != nil {
			log.Fatalln("Failed :", err.Error())
			return
		}
	case View:
		log.Fatalln("Failed  no route found:")
	}
}

func add_route(task *grpcc.Create_Task, connect *grpc.ClientConn) error {

	if reflect.DeepEqual(task, &grpcc.Create_Task{}) {
		return errors.New("empty task")
	}

	webtask := grpcc.NewWebTask_Object(connect)
	if status, _ := webtask.Add(context.Background(), task); reflect.DeepEqual(status, grpcc.Status{Result: grpcc.Result_Error}) {
		return errors.New("task is not created in memory ... :(")
	}

	if reflect.DeepEqual(task.TaskName, " ") {
		return errors.New("empty task name  .... :(")
	}

	if reflect.DeepEqual(task.Deadline, " ") {
		return errors.New("no deadline specified .... :(")
	}

	if task.GetLevelWork() < 0 && task.GetLevelWork() > 3 {
		return errors.New("priorty level is too high or too low .... :( ")
	}

	if reflect.DeepEqual(task.Id, " ") {
		return errors.New("id is not specified .... :( ")
	}

	if reflect.DeepEqual(task.Uid, " ") {
		return errors.New("uid is not specified .... :( ")
	}

	task.Uid = "_:" + task.TaskName

	log.Println("Data processing completed .... store in memory for future use .... :)")
	ngraph := graph.NewNGraph(task)

	status := ngraph.Create_Schema(task, context.Background())

	if reflect.DeepEqual(&status, &grpcc.Status{Result: grpcc.Result_Error}) {
		log.Fatalln("Failed to create schema", status)
		return errors.New("task is not created .... :(")
	}

	log.Println("Task created in memory ..... :)", status)

	task_created := &grpcc.Create_Task{TaskName: task.TaskName, Deadline: task.Deadline, LevelWork: grpcc.Priority(task.LevelWork), Id: task.GetId(), Uid: GetTask().TaskName}
	ngraph = graph.NewNGraph(task_created)
	status = ngraph.View_Schema(&grpcc.View_Task{ViewTask: task_created, Uid: GetTask().TaskName}, context.Background())

	if reflect.DeepEqual(status, grpcc.Status{Result: grpcc.Result_Error}) {
		log.Fatalln("Failed to View schema", status)
		return errors.New("Stack Empty ..... :(")
	}

	log.Println("Task Details ready to shared ..... :)")
	list := make([]*grpcc.View_Task, 0)
	list = append(list, &grpcc.View_Task{ViewTask: task_created})

	status, _ = webtask.List(context.Background(), list[0])
	if !reflect.DeepEqual(status, grpcc.Status{Result: grpcc.Result_Error}) {
		log.Println(" Hurrah   ..... :)")
	}

	return nil
}

func delete_route(t *grpcc.Create_Task, client *grpc.ClientConn) error {

	ngraph := graph.NewNGraph(&grpcc.Create_Task{TaskName: t.TaskName, Deadline: GetTask().Deadline, LevelWork: grpcc.Priority(GetTask().Level), Uid: "_:" + t.TaskName})
	status := ngraph.Delete_Schema(&grpcc.Delete_Task{TaskName: t.TaskName, Uid: GetTask().TaskName}, context.Background())

	if reflect.DeepEqual(status, &grpcc.Status{Result: grpcc.Result_Error}) {
		log.Fatalln("Failed to delete schema", status)
		return errors.New("no task found")
	}

	log.Println("Data defer from memory ..... :)")

	return nil
}

var Webroute route

func Set_Route(router route) { Webroute = router }
func Get_Route() route       { return Webroute }
