package delivery

import (
	"net/http"
	"synapsis-ecommerce/src/helper"
	"synapsis-ecommerce/src/helper/validator"

	"github.com/labstack/echo/v4"
)

func (d *delivery) registerHandler(c echo.Context) error {

	var user validator.UserRegister

	ctx := c.Request().Context()

	if err := c.Bind(&user); err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := c.Validate(user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = d.service.Register(ctx, user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "Succes Register", nil)

}

func (d *delivery) login(c echo.Context) error {

	var user validator.UserLogin

	ctx := c.Request().Context()

	if err := c.Bind(&user); err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := c.Validate(user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	token, err := d.service.Login(ctx, user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusNotFound, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "Successfully logged in. Token will expire after 24 hours", token)

}
