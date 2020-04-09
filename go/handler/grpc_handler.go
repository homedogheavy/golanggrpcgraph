package handler

import (
	"context"

	taskservice "../proto/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type grpcTask interface {
	GetTasks(ctx context.Context, taskQuery *taskservice.GetTaskRequest) (*taskservice.Tasks, error)
	PostTask(ctx context.Context, taskQuery *taskservice.PostTaskRequest) (*taskservice.Tasks, error)
	PutTask(ctx context.Context, taskQuery *taskservice.DeleteTaskRequest) (*taskservice.Tasks, error)
	DeleteTask(ctx context.Context, taskQuery *taskservice.DeleteTaskRequest) (*taskservice.Tasks, error)
}

type task struct {
	Id   int
	Name string
	Done bool
}

var grpcTasks = make([]*taskservice.Task, 0)

func NewGrpcTaskHandler() task {
	return task{}
}

func (r task) GetTasks(ctx context.Context, taskQuery *taskservice.GetTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &taskservice.Tasks{
		Tasks: grpcTasks,
	}, nil
}
func (r task) PostTask(ctx context.Context, taskQuery *taskservice.PostTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	task := &taskservice.Task{
		Id:   int64(len(grpcTasks) + 1),
		Name: taskQuery.Name,
		Done: false,
	}
	grpcTasks = append(grpcTasks, task)

	return &taskservice.Tasks{
		Tasks: grpcTasks,
	}, nil
}
func (r task) PutTask(ctx context.Context, taskQuery *taskservice.PutTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	for k, v := range grpcTasks {
		if v.Id == taskQuery.Id {
			grpcTasks[k].Done = taskQuery.Done
		}
	}
	return &taskservice.Tasks{
		Tasks: grpcTasks,
	}, nil

}
func (r task) DeleteTask(ctx context.Context, taskQuery *taskservice.DeleteTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	tasks := make([]*taskservice.Task, 0)
	for _, v := range grpcTasks {
		if int64(v.Id) != taskQuery.Id {
			tasks = append(tasks, v)
		}
	}
	grpcTasks = tasks
	return &taskservice.Tasks{
		Tasks: grpcTasks,
	}, nil
}
