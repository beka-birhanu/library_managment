package models

import (
	"fmt"
	"strings"
)

// Member represents a member with an ID, name, and a list of borrowed books.
type Member struct {
	id            int
	name          string
	borrowedBooks map[int]*Book
}

// NewMember creates a new Member instance with the given ID and name.
// Returns an error if any of the parameters are invalid.
func NewMember(id int, name string) (*Member, error) {
	name = strings.Trim(name, " ")

	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	return &Member{id: id, name: name, borrowedBooks: map[int]*Book{}}, nil
}

// ID returns the ID of the member.
func (m *Member) ID() int {
	return m.id
}

// Name returns the name of the member.
func (m *Member) Name() string {
	return m.name
}

// BorrowedBooks returns the list of books borrowed by the member.
func (m *Member) BorrowedBooks() []*Book {
	books := make([]*Book, 0)
	for _, book := range m.borrowedBooks {
		books = append(books, book)
	}
	return books
}

// BorrowedBooks returns the list of books borrowed by the member.
func (m *Member) Return(book *Book) error {
	_, borrowed := m.borrowedBooks[book.ID()]
	if !borrowed {
		return fmt.Errorf("book with ID %d was borrowed by a different member", book.ID())
	}

	book.status = StatusAvailable
	delete(m.borrowedBooks, book.ID())
	return nil
}

// Borrow allows the member to borrow a book if it is available.
// Returns an error if the book is not available or already borrowed.
func (m *Member) Borrow(book *Book) error {
	if book.Status() != StatusAvailable {
		return fmt.Errorf("book '%s' is not available for borrowing", book.Title())
	}
	book.status = StatusBorrowed
	m.borrowedBooks[book.ID()] = book
	return nil
}

