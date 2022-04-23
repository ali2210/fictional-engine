package graph

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	"github.com/go-redis/redis"

	"github.com/ali2210/fictional-engine/stubs/grpcc"
)

const (
	ADDRESS  string = "127.0.0.1:6379"
	DB       int    = 0
	PASSWORD string = ""
)

var redis_me *redis.Client

type Graph_Datasorage interface {
	Create_Schema(task *grpcc.Create_Task, ctx context.Context) *grpcc.Status
	Delete_Schema(task *grpcc.Delete_Task, ctx context.Context) *grpcc.Status
	View_Schema(task *grpcc.View_Task, ctx context.Context) *grpcc.Status
}

type NGraph struct {
	Task *grpcc.Create_Task
}

func NewNGraph(object *grpcc.Create_Task) Graph_Datasorage { return &NGraph{Task: object} }

func (g *NGraph) Create_Schema(task *grpcc.Create_Task, ctx context.Context) *grpcc.Status {

	if reflect.DeepEqual(task, grpcc.Create_Task{TaskName: "", Deadline: "", LevelWork: 0, Id: 0, Uid: "_:"}) {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     ADDRESS,
		Password: PASSWORD,
		DB:       DB,
	})

	redis_me = redisClient

	data, err := json.Marshal(task)
	if err != nil {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	err = redisClient.Set(task.Deadline, data, 0).Err()
	if err != nil {

		log.Println("error ... ", err)
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	log.Println("Task created in redis cluster: .... ")

	return &grpcc.Status{Result: grpcc.Result_Ok}
}

func (g *NGraph) Delete_Schema(task *grpcc.Delete_Task, ctx context.Context) *grpcc.Status {

	if reflect.DeepEqual(task.TaskName, " ") {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	status := g.View_Schema(&grpcc.View_Task{
		ViewTask: &grpcc.Create_Task{TaskName: g.Task.TaskName, Deadline: g.Task.Deadline, LevelWork: g.Task.LevelWork, Id: g.Task.Id},
	}, ctx)

	if reflect.DeepEqual(grpcc.Status{Result: grpcc.Result_Error}, status) {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	err := redis_me.Set(g.Task.Deadline, "", 0).Err()
	if err != nil {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	log.Println("Task deleted ..... :)")

	return &grpcc.Status{Result: grpcc.Result_Ok}

}

func (g *NGraph) View_Schema(task *grpcc.View_Task, ctx context.Context) *grpcc.Status {

	if reflect.DeepEqual(g.Task.TaskName, " ") {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	if reflect.DeepEqual(g.Task, &grpcc.Create_Task{}) {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	value, err := redis_me.Get(task.ViewTask.Deadline).Result()
	if err != nil {
		return &grpcc.Status{Result: grpcc.Result_Error}
	}

	log.Println("Here is your task ..... :)", value)
	return &grpcc.Status{Result: grpcc.Result_Ok}
}
