package helper

import (
	"github.com/labstack/echo/v4"
)

type response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteResponse(c echo.Context, status int, message interface{}, data interface{}) error {
	return c.JSON(status, response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
