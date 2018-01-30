package middleware

import (
	"github.com/Hexilee/rady"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"reflect"
)

type (
	JWTClaims struct {
		rady.Component
		Claims jwt.Claims
	}

	JWT struct {
		rady.Middleware
		SigningKey *string `value:"rady.jwt.signing_key"`
	}

	JWTWithConfig struct {
		rady.Middleware
		App           *rady.Application
		SigningKey    *string `value:"rady.jwt.signing_key"`
		TokenLookup   *string `value:"rady.jwt.token_lookup"`
		ContextKey    *string `value:"rady.jwt.context_key"`
		SigningMethod *string `value:"rady.jwt.signing_method"`
		AuthScheme    *string `value:"rady.jwt.auth_scheme"`
		Skipper       *string `value:"rady.middleware.logger.skipper" default:"GetJWTSkipper"`
		Claims        *string `value:"rady.jwt.claims" default:"GetJWTClaims"`
	}
)

var (
	JWTClaimsType = reflect.TypeOf(new(JWTClaims))
)

func GetClaims(App *rady.Application, Name string) jwt.Claims {
	BeanMap, ok := App.BeanMap[JWTClaimsType]
	if ok {
		claims, ok := BeanMap[Name]
		if ok {
			return (claims.Value.Addr().Interface()).(*JWTClaims).Claims
		}
	}
	return jwt.MapClaims{}
}

func (j *JWT) DefaultJWT(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.JWT(*j.SigningKey)(next)
}

func (j *JWTWithConfig) ConfigJWT(next rady.HandlerFunc) rady.HandlerFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(*j.SigningKey),
		TokenLookup:   *j.TokenLookup,
		ContextKey:    *j.ContextKey,
		SigningMethod: *j.SigningMethod,
		AuthScheme:    *j.AuthScheme,
		Skipper:       GetSkipper(j.App, *j.Skipper),
		Claims:        GetClaims(j.App, *j.Claims),
	})(next)
}
