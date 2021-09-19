package database

import (
	"github.com/DaffaBarin/simple-books-mvc/config"
	"github.com/DaffaBarin/simple-books-mvc/models"
)

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err:= config.DB.Find(&books).Error
	if err != nil{
		return []models.Book{}, err
	}

	return books, nil
}

func PostBook(book models.Book) (*models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}

func GetBookByID(ID int) (*models.Book, error) {
	var book models.Book

	err := config.DB.Where("id = ?", ID).First(&book).Error
	if err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}

func UpdateBook(book models.Book, ID int) (*models.Book, error) {
	err := config.DB.Where("id = ?", ID).Updates(book).Error

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func DeleteBook(ID int) (*models.Book, error) {
	var book models.Book

	err := config.DB.Where("id = ?", ID).Delete(&book).Error
	if err != nil {
		return &models.Book{}, err
	}

	return &book, nil
}