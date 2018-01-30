package middleware

import (
	"github.com/Hexilee/rady"
	"github.com/labstack/echo/middleware"
)

type Logger struct {
	rady.Middleware
}

func (logger *Logger)Log(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.Logger()(next)
}
