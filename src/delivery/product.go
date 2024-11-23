package delivery

import (
	"net/http"
	"synapsis-ecommerce/src/helper"

	"github.com/labstack/echo/v4"
)

func (d *delivery) getProducts(c echo.Context) error {
	ctx := c.Request().Context()

	products, err := d.service.GetProducts(ctx)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, "Success", products)
}

func (d *delivery) getProductsByCategory(c echo.Context) error {

	category := c.Param("category")
	ctx := c.Request().Context()

	products, err := d.service.GetProductByCategory(ctx, category)
	if err != nil {
		return helper.WriteResponse(c, http.StatusNotFound, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "Success", products)

}
