package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique;not null"`
	Author      string `json:"author" gorm:"not null"`
	Genre       string `json:"genre" gorm:"not null"`
	Description string `json:"description"`
	ISBN        string `json:"isbn" gorm:"not null"`
	Published   string `json:"published" gorm:"not null"`
	Publisher   string `json:"publisher" gorm:"not null"`
}

type BookList struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Published string `json:"published"`
	Publisher string `json:"publisher"`
}

type BookDetails struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	Published   string `json:"published"`
	Publisher   string `json:"publisher"`
}
