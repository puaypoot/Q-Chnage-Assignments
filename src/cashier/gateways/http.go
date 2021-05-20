package gateways

import (
	"q-change-assignment/src/cashier/dao"
	"q-change-assignment/src/cashier/services"

	echo "github.com/labstack/echo/v4"

	"net/http"
)

type HTTPGateway struct {
	CashierService services.ICashierService
}

// NewHTTPGateway NewHTTPGateway
func NewHTTPGateway(e *echo.Group, sv services.ICashierService) {
	handle := &HTTPGateway{
		CashierService: sv,
	}

	e.POST("/calculate-change-money", handle.CalculateChangeMoney)
	e.POST("/add-coins", handle.AddCoins)
}

func (h *HTTPGateway) CalculateChangeMoney(c echo.Context) (err error) {
	ctx := c.Request().Context()
	r := new(dao.ChangeMoneyRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string][]string{
			"errors": {err.Error()},
		})
	}

	res, err := h.CashierService.ChangeMoney(ctx, r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string][]string{
			"errors": {err.Error()},
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *HTTPGateway) AddCoins(c echo.Context) (err error) {
	ctx := c.Request().Context()
	r := new(dao.AddCoinRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string][]string{
			"errors": {err.Error()},
		})
	}

	res, errs := h.CashierService.AddCoins(ctx, r)
	if len(errs) > 0 {
		errMSGs := make([]string, 0, len(errs))
		for _, e := range errs {
			errMSGs = append(errMSGs, e.Error())
		}
		return c.JSON(http.StatusUnprocessableEntity, map[string][]string{
			"errors": errMSGs,
		})
	}

	return c.JSON(http.StatusOK, res)
}
