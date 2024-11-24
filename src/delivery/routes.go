package delivery

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (d *delivery) Auth(e *echo.Group) {

	e.POST("/register", d.registerHandler)
	e.POST("/login", d.login)
}

func (d *delivery) User(e *echo.Group, configJWT echojwt.Config) {
	e.Use(echojwt.WithConfig(configJWT))
	e.GET("/cart", d.getCartByUserId)
	e.POST("/cart", d.AddProductToCartUser)

}

func (d *delivery) Product(e *echo.Group, configJWT echojwt.Config) {
	e.Use(echojwt.WithConfig(configJWT))
	e.GET("", d.getProducts)
	e.GET("/:category", d.getProductsByCategory)
}

func (d *delivery) Routes(e *echo.Echo, configJWT echojwt.Config) {
	d.Auth(e.Group("auth"))
	d.Product(e.Group("products"), configJWT)
	d.User(e.Group("user"), configJWT)

}
