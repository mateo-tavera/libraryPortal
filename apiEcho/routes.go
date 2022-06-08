package apiecho

import (
	"github.com/labstack/echo"
)

func (a *API) RegisterRoutes(e *echo.Echo) {

	public := e.Group("")
	protected := e.Group("")

	//r.Use(RequestIdMiddleware)             //It means that any router uses this middleware
	//protected.Use(AuthorizationMiddleware) //Only the routes 'protected' uses this middleware

	public.GET("/books", a.GetBooks)
	public.GET("/book/:id", a.GetBook)

	protected.POST("/books", a.CreateBook)

}
