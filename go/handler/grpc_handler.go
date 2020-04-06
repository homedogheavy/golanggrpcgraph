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

func NewGrpcTaskHandler() task {
	return task{}
}

func (r task) GetTasks(ctx context.Context, taskQuery *taskservice.GetTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &taskservice.Tasks{
		Tasks: make([]*taskservice.Task, 0),
	}, nil
}
func (r task) PostTask(ctx context.Context, taskQuery *taskservice.PostTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &taskservice.Tasks{
		Tasks: make([]*taskservice.Task, 0),
	}, nil
}
func (r task) PutTask(ctx context.Context, taskQuery *taskservice.PutTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &taskservice.Tasks{
		Tasks: make([]*taskservice.Task, 0),
	}, nil
}
func (r task) DeleteTask(ctx context.Context, taskQuery *taskservice.DeleteTaskRequest) (*taskservice.Tasks, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

	return &taskservice.Tasks{
		Tasks: make([]*taskservice.Task, 0),
	}, nil
}
