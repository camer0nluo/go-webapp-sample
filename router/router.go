package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ybkuroki/go-webapp-sample/controller"
)

// Init is
func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders: []string{
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		MaxAge: 86400,
	}))
	e.Use(middleware.Recover())

	api := e.Group("/api")
	{
		book := api.Group("/book")
		{
			book.GET("/list", controller.GetBookList())
		}
	}

	return e
}