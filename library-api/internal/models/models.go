package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	ISBN      string         `json:"isbn" gorm:"unique"`
	Available bool           `json:"available" gorm:"default:true"`
	Authors   []Author       `json:"authors,omitempty" gorm:"many2many:book_authors;"`
	AuthorIDs []uint         `json:"author_ids,omitempty" gorm:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Author struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Bio       string         `json:"bio"`
	Books     []Book         `json:"books,omitempty" gorm:"many2many:book_authors;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Loan struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	BookID     uint       `json:"book_id" gorm:"not null"`
	Book       Book       `json:"book" gorm:"foreignKey:BookID"`
	UserName   string     `json:"user_name" gorm:"not null"`
	LoanDate   time.Time  `json:"loan_date"`
	ReturnDate *time.Time `json:"return_date"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
