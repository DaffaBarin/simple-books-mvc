package controllers

import (
	"time"
	"github.com/DaffaBarin/simple-books-mvc/models"
)

type RequestBook struct {
	Title     string    `json:"title"`
	ISBN	  string    `json:"isbn"`
	Writer	  string    `json:"writer"`
}

func (req *RequestBook) toModel() models.Book {
	return models.Book{
		Title:   req.Title,
		ISBN:  req.ISBN,
		Writer: req.Writer,
	}
}

type ResponseBook struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint    `json:"id"`
	Title     string    `json:"title"`
	ISBN	  string    `json:"isbn"`
	Writer	  string    `json:"writer"`
}

func newResponse(mdlBooks models.Book) ResponseBook {
	return ResponseBook{
		CreatedAt: mdlBooks.CreatedAt,
		UpdatedAt: mdlBooks.UpdatedAt,
		ID:        mdlBooks.ID,
		Title:     mdlBooks.Title,
		ISBN:     mdlBooks.ISBN,
		Writer:    mdlBooks.Writer,
	}
}

func newResponseArray(mdlBooks []models.Book) []ResponseBook {
	var resp []ResponseBook

	for _, value := range mdlBooks {
		resp = append(resp, newResponse(value))
	}

	return resp
}