package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID      uint `gorm:"primaryKey";type:autoIncrement`
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	ISBN  string `gorm:"type:varchar(255);not null" json:"isbn"`
	Writer string `gorm:"type:varchar(255);not null" json:"writer"`
}