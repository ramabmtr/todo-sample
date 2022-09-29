package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	RespStatusSuccess = "success"
	RespStatusError   = "error"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func BuildJSONResponse(c echo.Context, code int, msg string, data any, dataWrapper ...string) error {
	resp := Response{
		Status:  RespStatusSuccess,
		Message: msg,
		Data:    data,
	}

	if code > 399 {
		resp.Status = RespStatusError
	}

	for _, wrapper := range dataWrapper {
		data = map[string]any{
			wrapper: data,
		}
	}

	resp.Data = data

	return c.JSON(code, resp)
}

func ResponseOK(c echo.Context) error {
	return ResponseOKWithData(c, nil)
}

func ResponseOKWithData(c echo.Context, data any, dataWrapper ...string) error {
	return BuildJSONResponse(c, http.StatusOK, "", data, dataWrapper...)
}

func ResponseSuccess(c echo.Context, code int) error {
	return ResponseSuccessWithData(c, code, nil)
}

func ResponseSuccessWithData(c echo.Context, code int, data any, dataWrapper ...string) error {
	return BuildJSONResponse(c, code, "", data, dataWrapper...)
}
