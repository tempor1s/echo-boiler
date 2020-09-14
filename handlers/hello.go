package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Hello(c echo.Context) (err error) {
	resp := map[string]string{"msg": "Hello!"}

	return c.JSON(http.StatusOK, resp)
}
