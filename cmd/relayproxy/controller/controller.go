package controller

import "github.com/labstack/echo/v4"

// Controller is the interface used by all controller struct
type Controller interface {
	Handler(c echo.Context) error
}

type OFREPController interface {
	Handler(c echo.Context) error
	OFREPHandler(c echo.Context) error
}
