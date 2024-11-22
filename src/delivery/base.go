package delivery

import (
	"synapsis-ecommerce/src/service"

	"github.com/labstack/echo/v4"
)

type Delivery interface {
	RegisterHandler(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Auth(e *echo.Group)
}

type delivery struct {
	service service.Service
}

func NewDelivery(service service.Service) Delivery {
	return &delivery{
		service: service,
	}
}
