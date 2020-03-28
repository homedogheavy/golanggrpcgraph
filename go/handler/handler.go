package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type Task struct {
	Id   int
	Name string
	Done bool
}

var Tasks = make([]Task, 0)

func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, Tasks)
	}
}

func PostTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &struct {
			Name string
		}{}
		if err := c.Bind(req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		task := Task{
			Id:   len(Tasks) + 1,
			Name: req.Name,
			Done: false,
		}
		Tasks = append(Tasks, task)
		return c.NoContent(http.StatusOK)
	}
}
func PutTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &struct {
			Id   int
			Done bool
		}{}
		if err := c.Bind(req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		for k, v := range Tasks {
			if v.Id == req.Id {
				Tasks[k].Done = req.Done
			}
		}
		return c.NoContent(http.StatusOK)
	}
}
func DeleteTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &struct {
			Id int
		}{}
		if err := c.Bind(req); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		tasks := make([]Task, 0)
		for _, v := range Tasks {
			if v.Id != req.Id {
				tasks = append(tasks, v)
			}
		}
		Tasks = tasks
		return c.NoContent(http.StatusOK)
	}
}
