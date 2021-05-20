package gateways

import (
	"q-change-assignment/src/xyz/services"

	echo "github.com/labstack/echo/v4"

	"net/http"
)

type HTTPGateway struct {
	XYZService services.IXYZService
}

// NewHTTPGateway NewHTTPGateway
func NewHTTPGateway(e *echo.Group, sv services.IXYZService) {
	handle := &HTTPGateway{
		XYZService: sv,
	}

	e.GET("/", handle.FindXYZ)
}

func (h *HTTPGateway) FindXYZ(c echo.Context) (err error) {
	res, err := h.XYZService.FindXYZ(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string][]string{
			"errors": {err.Error()},
		})
	}
	return c.JSON(http.StatusOK, res)
}
