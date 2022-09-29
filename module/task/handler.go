package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/todo-sample/util"
)

type handler struct {
	useCase UseCaseIFace
}

func NewHandler(useCase UseCaseIFace) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = c.Validate(&req)
	if err != nil {
		return err
	}

	resp, err := h.useCase.Create(ctx, &req)
	if err != nil {
		return err
	}

	return util.ResponseOKWithData(c, resp)
}

func (h *handler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	limitStr := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitStr)

	pageStr := c.QueryParam("offset")
	page, _ := strconv.Atoi(pageStr)

	completeStatusStr := c.QueryParam("completeStatus")
	completeStatusInt, _ := strconv.Atoi(completeStatusStr)

	var completeStatus CompleteStatus
	switch CompleteStatus(completeStatusInt) {
	case All, Complete, NotComplete:
		completeStatus = CompleteStatus(completeStatusInt)
	default:
		completeStatus = All
	}

	resp, total, err := h.useCase.Get(ctx, limit, page, completeStatus)
	if err != nil {
		return err
	}

	return util.ResponseOKWithData(c, map[string]interface{}{
		"data": resp,
		"pagination": map[string]interface{}{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (h *handler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is missing")
	}

	resp, err := h.useCase.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return util.ResponseOKWithData(c, resp)
}

func (h *handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is missing")
	}

	var req UpdateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = c.Validate(&req)
	if err != nil {
		return err
	}

	resp, err := h.useCase.UpdateByID(ctx, id, &req)
	if err != nil {
		return err
	}

	return util.ResponseOKWithData(c, resp)
}

func (h *handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is missing")
	}

	err := h.useCase.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return util.ResponseOK(c)
}
