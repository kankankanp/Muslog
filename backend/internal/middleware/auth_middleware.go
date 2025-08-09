package middleware

import (
	"net/http"
<<<<<<< HEAD
	"strings"
=======
>>>>>>> develop

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type AuthMiddlewareConfig struct {
<<<<<<< HEAD
	Skipper   echoMiddleware.Skipper
	JWTSecret string
=======
	Skipper echoMiddleware.Skipper
>>>>>>> develop
}

func AuthMiddleware(config AuthMiddlewareConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = echoMiddleware.DefaultSkipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

<<<<<<< HEAD
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Missing or invalid token"})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return []byte(config.JWTSecret), nil
=======
			cookie, err := c.Cookie("access_token")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Missing or invalid token"})
			}

			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return []byte("secret"), nil
>>>>>>> develop
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid or expired token"})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token claims"})
			}

			c.Set("user", claims)

			return next(c)
		}
	}
<<<<<<< HEAD
}
=======
}
>>>>>>> develop
