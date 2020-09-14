package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Hello(c echo.Context) (err error) {
	resp := map[string]string{"msg": "Hello!"}

	return c.JSON(http.StatusOK, resp)
}
