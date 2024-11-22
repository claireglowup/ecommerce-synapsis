package src

import (
	"synapsis-ecommerce/config/util"
	"synapsis-ecommerce/src/helper/validator"

	validatorEngine "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Run()
}

type server struct {
	httpServer *echo.Echo
	config     util.Config
}

func InitServer(config util.Config) Server {

	e := echo.New()
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		config:     config,
		httpServer: e,
	}

}


func (s *server) Run() {}