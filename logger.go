package middleware

import (
	"github.com/Hexilee/rady"
	"github.com/labstack/echo/middleware"
	"io"
	"os"
	"reflect"
)

type (
	Logger struct {
		rady.Middleware
	}

	LoggerOutput struct {
		rady.Component
		OutPut io.Writer
	}

	LogWithConfig struct {
		rady.Middleware
		App     *rady.Application
		Format  *string `value:"rady.middleware.logger.format"`
		Skipper *string `value:"rady.middleware.logger.skipper" default:"GetLoggerSkipper"`
		Output  *string `value:"rady.middleware.logger.output" default:"GetLoggerOutput"`
	}
)

var (
	OutputType = reflect.TypeOf(new(LoggerOutput))
)

func GetOutput(App *rady.Application, Name string) io.Writer {
	BeanMap, ok := App.BeanMap[OutputType]
	if ok {
		output, ok := BeanMap[Name]
		if ok {
			return (output.Value.Addr().Interface()).(*LoggerOutput).OutPut
		}
	}
	return os.Stdout
}

func (logger *Logger) DefaultLogger(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.Logger()(next)
}

func (logger *LogWithConfig) ConfigLogger(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  *logger.Format,
		Skipper: GetSkipper(logger.App, *logger.Skipper),
		Output:  GetOutput(logger.App, *logger.Output),
	})(next)
}
