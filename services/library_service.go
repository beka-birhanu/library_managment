package services

import (
	"fmt"

	"github.com/beka-birhanu/library_managment/common"
	"github.com/beka-birhanu/library_managment/models"
)

// Library represents a library with books and members.
type Library struct {
	availableBooks map[int]*models.Book
	members        map[int]*models.Member
	borrowedBooks  map[int]*models.Book
}

// Making sure Library implements common.LibraryManager
var _ common.LibraryManager = &Library{}

// NewLibrary creates a new Library instance.
func NewLibrary() *Library {
	return &Library{
		availableBooks: make(map[int]*models.Book),
		members:        make(map[int]*models.Member),
		borrowedBooks:  make(map[int]*models.Book),
	}
}

// AddBook adds a new book to the library.
func (l *Library) AddBook(book *models.Book) error {
	if _, ok := l.availableBooks[book.ID()]; ok {
		return fmt.Errorf("a book with id %d already exists", book.ID())
	}
	l.availableBooks[book.ID()] = book
	return nil
}

// RemoveBook removes a book from the library by its ID.
func (l *Library) RemoveBook(bookID int) error {
	_, available := l.availableBooks[bookID]
	if !available {
		return fmt.Errorf("the book is borrowed, it has to be returned before to delete")
	}
	delete(l.availableBooks, bookID)
	return nil
}

// BorrowBook allows a member to borrow a book by its ID.
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.availableBooks[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d does not exist", bookID)
	}
	if book.Status() != models.StatusAvailable {
		return fmt.Errorf("book with ID %d is not available", bookID)
	}
	member, exists := l.members[memberID]
	if !exists {
		return fmt.Errorf("member with ID %d does not exist", memberID)
	}

	err := member.Borrow(book)
	if err != nil {
		return err
	}

	delete(l.availableBooks, bookID)
	l.borrowedBooks[book.ID()] = book
	return nil
}

// ReturnBook allows a member to return a borrowed book by its ID.
func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := l.members[memberID]
	if !exists {
		return fmt.Errorf("member with ID %d does not exist", memberID)
	}

	book, borrowed := l.borrowedBooks[bookID]
	if !borrowed {
		return fmt.Errorf("book with ID %d was not borrowed", bookID)
	}

	err := member.Return(book)
	if err != nil {
		return err
	}

	delete(l.borrowedBooks, bookID)
	l.availableBooks[book.ID()] = book
	return nil
}

// ListAvailableBooks returns a list of books that are currently available.
func (l *Library) ListAvailableBooks() []*models.Book {
	availableBooks := make([]*models.Book, 0)
	for _, book := range l.availableBooks {
		availableBooks = append(availableBooks, book)
	}
	return availableBooks
}

// ListBorrowedBooks returns a list of books borrowed by a specific member.
func (l *Library) ListBorrowedBooks(memberID int) []*models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks()
}
