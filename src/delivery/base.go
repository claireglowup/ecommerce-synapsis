package delivery

import (
	"synapsis-ecommerce/src/service"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Delivery interface {
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
