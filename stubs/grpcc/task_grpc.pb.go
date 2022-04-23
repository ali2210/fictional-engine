package grpcc

import (
	"context"
	"errors"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

type Task_Feature_Client interface {
	Add(ctx context.Context, t *Create_Task, call ...grpc.CallOption) (*Status, error)
	Remove(ctx context.Context, t *Delete_Task, call ...grpc.CallOption) (*Status, error)
	List(ctx context.Context, t *View_Task, call ...grpc.CallOption) (*Status, error)
}

type task_feature_client struct {
	Connection grpc.ClientConnInterface
}

func NewWebTask_Object(con grpc.ClientConnInterface) Task_Feature_Client {
	return &task_feature_client{Connection: con}
}

func (c *task_feature_client) Add(ctx context.Context, t *Create_Task, call ...grpc.CallOption) (*Status, error) {

	out := new(Status)
	err := c.Connection.Invoke(ctx, "/stubs/grpcc.Task/Add", t, out, call...)
	if err != nil {
		return &Status{Result: Result_Error}, err
	}
	return &Status{Result: Result_Ok}, nil
}

func (c *task_feature_client) Remove(ctx context.Context, t *Delete_Task, call ...grpc.CallOption) (*Status, error) {

	out := new(Status)
	err := c.Connection.Invoke(ctx, "/stubs/grpcc.Task/Remove", t, out, call...)
	if err != nil {
		return &Status{Result: Result_Error}, err
	}
	return &Status{Result: Result_Ok}, nil
}

func (c *task_feature_client) List(ctx context.Context, t *View_Task, call ...grpc.CallOption) (*Status, error) {

	out := new(Status)
	err := c.Connection.Invoke(ctx, "/stubs/grpcc.Task/List", t, out, call...)
	if err != nil {
		return &Status{Result: Result_Error}, err
	}
	return &Status{Result: Result_Ok}, nil

}

type Plan int64

type TaskRouteServer interface {
	Add(ctx context.Context, t *Create_Task) (*Status, error)
	Remove(ctx context.Context, t *Delete_Task) (*Status, error)
	List(ctx context.Context, t *View_Task) (*Status, error)
	mustEmbedUnimplementedTaskService()
}

func (imp UnimplementedTaskPlan) Add(ctx context.Context, t *Create_Task, call ...grpc.CallOption) (*Status, error) {
	return &Status{Result: Result_Error}, status.Errorf(codes.Unimplemented, "add method is not implemented")
}

func (imp UnimplementedTaskPlan) Remove(ctx context.Context, t *Delete_Task, call ...grpc.CallOption) (*Status, error) {
	return &Status{Result: Result_Error}, status.Errorf(codes.Unimplemented, "delete method is not implemented")
}

func (impl UnimplementedTaskPlan) List(ctx context.Context, t *View_Task, call ...grpc.CallOption) (*Status, error) {
	return &Status{Result: Result_Error}, status.Errorf(codes.Unimplemented, "list method is not implemented")
}

func (imp UnimplementedTaskPlan) mustEmbedUnimplementedTaskService() {}

type UnimplementedTaskPlan struct{}

func RegisterService(registrar grpc.ServiceRegistrar, sv TaskRouteServer) {
	registrar.RegisterService(&Task_Service_Desc, sv)
}

type UnSafeTaskService interface {
	mustEmbedUnimplementedTaskService()
}

func _TaskService_AddRoute(server interface{}, ctx context.Context, desc func(interface{}) error, intercept grpc.UnaryServerInterceptor) (interface{}, error) {

	service_add := new(Create_Task)
	err := desc(service_add)
	if err != nil {
		return nil, err
	}

	if intercept == nil {
		server.(TaskRouteServer).Add(ctx, service_add)
	}

	serverInfo := &grpc.UnaryServerInfo{
		Server:     server,
		FullMethod: "/stubs/grpcc.Task/Add",
	}

	handle := func(ctx context.Context, request interface{}) (interface{}, error) {

		istatus, _ := server.(TaskRouteServer).Add(ctx, request.(*Create_Task))
		if reflect.DeepEqual(istatus, Status{Result: Result_Error}) {
			return istatus, errors.New("task is not created")
		}
		return server.(TaskRouteServer).Add(ctx, service_add)
	}

	return intercept(ctx, service_add, serverInfo, handle)
}

func _TaskService_ListRoute(server interface{}, ctx context.Context, desc func(interface{}) error, intercept grpc.UnaryServerInterceptor) (interface{}, error) {

	service_list := new(View_Task)
	err := desc(service_list)
	if err != nil {
		return nil, err
	}

	if intercept == nil {
		server.(TaskRouteServer).List(ctx, service_list)
	}

	serverInfo := &grpc.UnaryServerInfo{
		Server:     server,
		FullMethod: "/stubs/grpcc.Task/List",
	}

	handle := func(ctx context.Context, request interface{}) (interface{}, error) {

		istatus, _ := server.(TaskRouteServer).List(ctx, request.(*View_Task))
		if reflect.DeepEqual(istatus, Status{Result: Result_Error}) {
			return istatus, errors.New("no task returned")
		}
		return server.(TaskRouteServer).List(ctx, request.(*View_Task))
	}

	return intercept(ctx, service_list, serverInfo, handle)

}

func _TaskService_DelRoute(server interface{}, ctx context.Context, desc func(interface{}) error, intercept grpc.UnaryServerInterceptor) (interface{}, error) {

	service_del := new(Delete_Task)
	err := desc(service_del)
	if err != nil {
		return nil, err
	}

	if intercept == nil {
		server.(TaskRouteServer).Remove(ctx, service_del)
	}

	serverInfo := &grpc.UnaryServerInfo{
		Server:     server,
		FullMethod: "/stubs/grpcc.Task/Delete",
	}

	handle := func(ctx context.Context, request interface{}) (interface{}, error) {

		istatus, _ := server.(TaskRouteServer).List(ctx, request.(*View_Task))
		if reflect.DeepEqual(istatus, Status{Result: Result_Error}) {
			return istatus, errors.New("no task returned")
		}
		return istatus, nil
	}

	return intercept(ctx, service_del, serverInfo, handle)
}

var Task_Service_Desc = grpc.ServiceDesc{
	ServiceName: "grpcc.Task",
	HandlerType: (*TaskRouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TaskService_AddRoute,
		},
		{
			MethodName: "Delete",
			Handler:    _TaskService_DelRoute,
		},
		{
			MethodName: "List",
			Handler:    _TaskService_ListRoute,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ali2210/fictional-engine/stubs/grpcc/task.proto",
}
