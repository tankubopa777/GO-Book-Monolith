package handlers

import "github.com/labstack/echo/v4"

type UserHandler interface {
 DetectUser(c echo.Context) error
}