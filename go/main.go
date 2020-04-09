package main

import (
	"net/http"

	"./handler"
	taskservice "./proto/task"
	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	e.Use(gRPCMiddleware(grpc.NewServer()))

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// ルーティング
	e.GET("/tasks", handler.GetTasks())
	e.POST("/tasks", handler.PostTasks())
	e.PUT("/tasks", handler.PutTasks())
	e.DELETE("/tasks", handler.DeleteTasks())

	// サーバー起動
	e.Start(":1323")
}
func gRPCMiddleware(grpcServer *grpc.Server) func(next echo.HandlerFunc) echo.HandlerFunc {

	grpcTaskHandler := handler.NewGrpcTaskHandler()
	taskservice.RegisterTaskServiceServer(grpcServer, grpcTaskHandler)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			wrappedGrpc := grpcweb.WrapServer(grpcServer)
			req := c.Request()
			resp := c.Response()
			if wrappedGrpc.IsGrpcWebRequest(req) {
				wrappedGrpc.ServeHTTP(resp, req)
			}
			return next(c)
		}
	}
}
