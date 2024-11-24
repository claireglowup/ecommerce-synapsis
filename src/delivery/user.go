package delivery

import (
	"net/http"
	"synapsis-ecommerce/src/helper"
	"synapsis-ecommerce/src/helper/validator"

	"github.com/labstack/echo/v4"
)

func (d *delivery) getCartByUserId(c echo.Context) error {

	authHeader := c.Request().Header.Get("Authorization")
	ctx := c.Request().Context()
	cartUser, err := d.service.GetCartByUserId(ctx, authHeader)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, "Success", cartUser)

}

func (d *delivery) AddProductToCartUser(c echo.Context) error {

	authHeader := c.Request().Header.Get("Authorization")

	ctx := c.Request().Context()

	var arg validator.AddProductCartUser

	if err := c.Bind(&arg); err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err := d.service.AddProductToCartTx(ctx, authHeader, arg)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, 200, "Success add product to cart", nil)

}
