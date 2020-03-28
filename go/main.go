package main

import (
	"net/http"

	"./handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

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
