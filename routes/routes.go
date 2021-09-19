package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/DaffaBarin/simple-books-mvc/controllers"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	e.GET("/",controllers.GreetingMessage)
	e.GET("/book", controllers.GetAllBooksController)
	e.POST("/book", controllers.PostBookController)
	e.GET("/book/:id", controllers.GetBookByIDController)
	e.PUT("/book/:id", controllers.UpdateBookController)
	e.DELETE("/book/:id", controllers.DeleteBookController)

	return e
}