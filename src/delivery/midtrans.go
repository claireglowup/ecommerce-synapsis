package delivery

import (
	"net/http"
	"synapsis-ecommerce/src/helper"

	"github.com/labstack/echo/v4"
)

func (d *delivery) MidtransPayment(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")

	ctx := c.Request().Context()

	midtransResponse, err := d.service.MidtransPayment(ctx, authHeader)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, 200, "Success", midtransResponse)
}
