package common

import (
	"github.com/beka-birhanu/library_managment/models"
)

// LibraryManager defines the interface for managing a library's books and member operations.
type LibraryManager interface {
	// AddBook adds a new book to the library.
	// Returns an error if there is ID conflict.
	AddBook(book *models.Book) error

	// RemoveBook removes a book from the library by its ID.
	RemoveBook(bookID int) error

	// BorrowBook allows a member to borrow a book by its ID.
	// Returns an error if the book cannot be borrowed (e.g., if it is not available).
	BorrowBook(bookID int, memberID int) error

	// ReturnBook allows a member to return a borrowed book by its ID.
	// Returns an error if the book cannot be returned (e.g., if it was not borrowed by the member).
	ReturnBook(bookID int, memberID int) error

	// ListAvailableBooks returns a list of books that are currently available.
	ListAvailableBooks() []*models.Book

	// ListBorrowedBooks returns a list of books borrowed by a specific member.
	ListBorrowedBooks(memberID int) []*models.Book
}
