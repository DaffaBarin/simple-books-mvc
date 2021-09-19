package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/DaffaBarin/simple-books-mvc/lib/database"
	"strconv"
)

func GreetingMessage(echoContext echo.Context) error {

	return echoContext.JSON(http.StatusOK, "Selamat datang di API buku buatan Daffa ges")
}

func GetAllBooksController(echoContext echo.Context) error {

	books, err := database.GetAllBooks()
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, books)
}

func PostBookController(echoContext echo.Context) error {

	var bookReq RequestBook
	echoContext.Bind(&bookReq)
	result, err := database.PostBook(bookReq.toModel())
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   newResponse(*result),
	})
}

func GetBookByIDController(echoContext echo.Context) error {
	ID := echoContext.Param("id")
	intID, _ := strconv.Atoi(ID)
	book, e := database.GetBookByID(intID)
	if e != nil {
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, book)
}

func UpdateBookController(echoContext echo.Context) error {
	ID := echoContext.Param("id")
	intID, _ := strconv.Atoi(ID)

	var bookReq RequestBook
	echoContext.Bind(&bookReq)
	result, err := database.UpdateBook(bookReq.toModel(),intID)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	result, _ = database.GetBookByID(intID)
	return echoContext.JSON(http.StatusOK, result)
}

func DeleteBookController(echoContext echo.Context) error {
	ID := echoContext.Param("id")
	intID, _ := strconv.Atoi(ID)
	_, e := database.DeleteBook(intID)
	if e != nil {
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}