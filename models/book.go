package models

import (
	"fmt"
	"strings"
)

// Define constants for the book status.
const (
	StatusAvailable = "available"
	StatusBorrowed  = "borrowed"
)

// Book represents a book with an ID, title, author, and status.
type Book struct {
	id     int
	title  string
	author string
	status string
}

// NewBook creates a new Book instance and validates parameters.
// Returns an error if any of the parameters are invalid.
func NewBook(id int, title, author, status string) (*Book, error) {
	title = strings.Trim(title, " ")
	author = strings.Trim(author, " ")
	status = strings.Trim(status, " ")

	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}
	if author == "" {
		return nil, fmt.Errorf("author cannot be empty")
	}
	if err := validateStatus(status); err != nil {
		return nil, err
	}

	return &Book{id: id, title: title, author: author, status: status}, nil
}

// validateStatus checks if the status is either StatusAvailable or StatusBorrowed.
// Returns an error if the status is invalid.
func validateStatus(status string) error {
	if status != StatusAvailable && status != StatusBorrowed {
		return fmt.Errorf("status must be either '%s' or '%s'", StatusAvailable, StatusBorrowed)
	}
	return nil
}

// ID returns the ID of the book.
func (b *Book) ID() int {
	return b.id
}

// Title returns the title of the book.
func (b *Book) Title() string {
	return b.title
}

// Author returns the author of the book.
func (b *Book) Author() string {
	return b.author
}

// Status returns the status of the book.
func (b *Book) Status() string {
	return b.status
}

