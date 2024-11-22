package delivery

import (
	"net/http"
	"synapsis-ecommerce/src/helper"
	"synapsis-ecommerce/src/helper/validator"

	"github.com/labstack/echo/v4"
)

func (d *delivery) RegisterHandler(c echo.Context) error {

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

func (d *delivery) Login(c echo.Context) error {

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

	return helper.WriteResponse(c, http.StatusOK, "Succesfully Login", token)

}

func (d *delivery) Logout(c echo.Context) error {

	return helper.WriteResponse(c, http.StatusOK, "you are logout", nil)

}
