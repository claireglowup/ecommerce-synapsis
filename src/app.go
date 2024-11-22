package src

import (
	"log"
	"synapsis-ecommerce/config/db"
	"synapsis-ecommerce/config/util"
	"synapsis-ecommerce/src/delivery"
	"synapsis-ecommerce/src/helper/validator"
	"synapsis-ecommerce/src/repository"
	"synapsis-ecommerce/src/service"

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

func InitServer(config *util.Config) Server {

	e := echo.New()
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		config:     *config,
		httpServer: e,
	}

}

func (s *server) Run() {

	pg := db.NewSQL(s.config.DBDriver, s.config.DatabaseURL)
	pg.InitPostgre()
	pg.RunDBMigration(s.config.MigrationURL)
	defer pg.DB.Close()

	repo := repository.New(pg.DB)
	service := service.NewService(repo)
	delivery := delivery.NewDelivery(service)
	delivery.Auth(s.httpServer.Group("auth"))
	// delivery.User(s.httpServer.Group("user"))

	if err := s.httpServer.Start(s.config.HTTPServerAddress); err != nil {
		log.Fatal(err.Error())
	}
}
