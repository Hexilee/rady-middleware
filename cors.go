package middleware

import (
	"github.com/Hexilee/rady"
	"github.com/labstack/echo/middleware"
)

type (
	CORS struct {
		rady.Middleware
	}

	CORSWithConfig struct {
		rady.Middleware
		App     *rady.Application
		Skipper *string `value:"rady.cors.skipper" default:"GetCORSSkipper"`
		// AllowOrigin defines a list of origins that may access the resource.
		// Optional. Default value *[]interface{}{"*"}.
		AllowOrigins *[]interface{} `value:"rady.cors.allow_origins"`

		// AllowMethods defines a list methods allowed when accessing the resource.
		// This is used in response to a preflight request.
		// Optional. Default value DefaultCORSConfig.AllowMethods.

		AllowMethods *[]interface{} `value:"rady.cors.allow_methods"`

		// AllowHeaders defines a list of request headers that can be used when
		// making the actual request. This in response to a preflight request.
		// Optional. Default value *[]interface{}{}.
		AllowHeaders *[]interface{} `value:"rady.cors.allow_headers"`

		// AllowCredentials indicates whether or not the response to the request
		// can be exposed when the credentials flag is true. When used as part of
		// a response to a preflight request, this indicates whether or not the
		// actual request can be made using credentials.
		// Optional. Default value false.
		AllowCredentials *bool `value:"rady.cors.allow_credentials"`

		// ExposeHeaders defines a whitelist headers that clients are allowed to
		// access.
		// Optional. Default value *[]interface{}{}.
		ExposeHeaders *[]interface{} `value:"rady.cors.expose_headers"`

		// MaxAge indicates how long (in seconds) the results of a preflight request
		// can be cached.
		// Optional. Default value 0.
		MaxAge *int `value:"rady.cors.max_age"`
	}
)

func InterfaceToString(interfaces *[]interface{}) []string {
	stringSlice := make([]string, 0)
	for _, value := range *interfaces {
		str, ok := value.(string)
		if !ok {
			continue
		}
		stringSlice = append(stringSlice, str)
	}
	return stringSlice
}

func (cors *CORS) DefaultCORS(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.CORS()(next)
}

func (cors *CORSWithConfig) ConfigCORS(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     InterfaceToString(cors.AllowOrigins),
		AllowMethods:     InterfaceToString(cors.AllowMethods),
		AllowHeaders:     InterfaceToString(cors.AllowHeaders),
		AllowCredentials: *cors.AllowCredentials,
		ExposeHeaders:    InterfaceToString(cors.ExposeHeaders),
		MaxAge:           *cors.MaxAge,
		Skipper:          GetSkipper(cors.App, *cors.Skipper),
	})(next)
}
