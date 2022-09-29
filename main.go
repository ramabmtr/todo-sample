package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickmn/go-cache"
	"github.com/ramabmtr/todo-sample/config"
	"github.com/ramabmtr/todo-sample/module/task"
	"github.com/ramabmtr/todo-sample/util"
)

type EchoValidator struct {
	v *validator.Validate
}

func (ev *EchoValidator) Validate(a any) error {
	if err := ev.v.Struct(a); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &EchoValidator{v: validator.New()}
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := http.StatusText(code)

		if e, ok := err.(*echo.HTTPError); ok {
			code = e.Code
			msg = fmt.Sprintf("%v", e.Message)
		}
		_ = util.BuildJSONResponse(c, code, msg, nil)
	}

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	taskRepo := task.NewRepo(cache.New(cache.NoExpiration, time.Hour))
	taskUseCase := task.NewUseCase(taskRepo)
	taskHandler := task.NewHandler(taskUseCase)
	taskGroup := e.Group("/task")
	taskHandler.Mount(taskGroup)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.App.Port)))
}
