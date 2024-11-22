package delivery

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (d *delivery) Auth(e *echo.Group) {

	e.POST("/register", d.RegisterHandler)
	e.POST("/login", d.Login)
	e.POST("/logout", d.Logout)
}

func (d *delivery) User(e *echo.Group) {

}

func (d *delivery) Product(e *echo.Group, configJWT echojwt.Config) {
	e.Use(echojwt.WithConfig(configJWT))
	e.GET("", d.GetProducts)
}

func (d *delivery) Routes(e *echo.Echo, configJWT echojwt.Config) {
	d.Auth(e.Group("auth"))
	d.Product(e.Group("products"), configJWT)
	d.User((e.Group("user")))

}
