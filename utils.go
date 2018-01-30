package middleware

import (
	"github.com/Hexilee/rady"
	"github.com/labstack/echo/middleware"
	"reflect"
)

var (
	SkipperType = reflect.TypeOf(new(RadySkipper))
)

type RadySkipper struct {
	rady.Component
	Skipper middleware.Skipper
}

func NewSkipper(Skipper middleware.Skipper) *RadySkipper {
	return &RadySkipper{Skipper: Skipper}
}

func GetSkipper(App *rady.Application, Name string) middleware.Skipper {
	BeanMap, ok := App.BeanMap[SkipperType]
	if ok {
		skipper, ok := BeanMap[Name]
		if ok {
			return (skipper.Value.Addr().Interface()).(*RadySkipper).Skipper
		}
	}
	return middleware.DefaultSkipper
}
