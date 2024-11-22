package delivery

import (
	"synapsis-ecommerce/src/service"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Delivery interface {
	RegisterHandler(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	Routes(e *echo.Echo, configJWT echojwt.Config)
}

type delivery struct {
	service service.Service
}

func NewDelivery(service service.Service) Delivery {
	return &delivery{
		service: service,
	}
}
