package delivery

import (
	"net/http"
	"synapsis-ecommerce/src/helper"

	"github.com/labstack/echo/v4"
)

func (d *delivery) GetProducts(c echo.Context) error {
	ctx := c.Request().Context()

	products, err := d.service.GetProducts(ctx)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, "success", products)
}
