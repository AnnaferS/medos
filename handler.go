package main

import (

	"github.com/labstack/echo"
	"net/http"
)

type handler struct{}

//Маршрут 1
//Создает пару токенов
func (h *handler) login(c echo.Context) error {
	guid := c.FormValue("GUID")

	tokens, err := generateTokenPair(guid)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tokens)
}
