package controllers

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/beka-birhanu/library_managment/common"
	"github.com/beka-birhanu/library_managment/models"
)

type Console struct {
	libraryService common.LibraryManager
	reader         io.Reader
	writer         io.Writer
}

func NewConsole(libraryService common.LibraryManager, reader io.Reader, writer io.Writer) *Console {
	return &Console{
		libraryService: libraryService,
		reader:         reader,
		writer:         writer,
	}
}

func (c *Console) Run() {
	scanner := bufio.NewScanner(c.reader)

	for {
		fmt.Fprintln(c.writer, "Library Management System")
		fmt.Fprintln(c.writer, "1. Add a new book")
		fmt.Fprintln(c.writer, "2. Remove an existing book")
		fmt.Fprintln(c.writer, "3. Borrow a book")
		fmt.Fprintln(c.writer, "4. Return a book")
		fmt.Fprintln(c.writer, "5. List all available books")
		fmt.Fprintln(c.writer, "6. List all borrowed books by a member")
		fmt.Fprintln(c.writer, "7. Exit")

		fmt.Fprint(c.writer, "Enter your choice: ")
		scanner.Scan()
		choiceStr := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Fprintln(c.writer, "Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			c.addBook(scanner)
		case 2:
			c.removeBook(scanner)
		case 3:
			c.borrowBook(scanner)
		case 4:
			c.returnBook(scanner)
		case 5:
			c.listAvailableBooks()
		case 6:
			c.listBorrowedBooks(scanner)
		case 7:
			fmt.Fprintln(c.writer, "Exiting...")
			return
		default:
			fmt.Fprintln(c.writer, "Invalid choice, please try again.")
		}
	}
}

func (c *Console) addBook(scanner *bufio.Scanner) {
	fmt.Fprint(c.writer, "Enter book ID: ")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid book ID. Please enter a number.")
		return
	}

	fmt.Fprint(c.writer, "Enter book title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	fmt.Fprint(c.writer, "Enter book author: ")
	scanner.Scan()
	author := strings.TrimSpace(scanner.Text())

	fmt.Fprint(c.writer, "Enter book status (available/borrowed): ")
	scanner.Scan()
	status := strings.TrimSpace(scanner.Text())

	book, err := models.NewBook(id, title, author, status)
	if err != nil {
		fmt.Fprintln(c.writer, "Error adding book:", err)
		return
	}

	err = c.libraryService.AddBook(book)
	if err != nil {
		fmt.Fprintln(c.writer, "Error adding Book! :$v", err.Error())
	} else {
		fmt.Fprintln(c.writer, "Book added successfully!")
	}
}

func (c *Console) removeBook(scanner *bufio.Scanner) {
	fmt.Fprint(c.writer, "Enter book ID to remove: ")
	scanner.Scan()
	idStr := strings.TrimSpace(scanner.Text())
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid book ID. Please enter a number.")
		return
	}

	err = c.libraryService.RemoveBook(id)
	if err != nil {
		fmt.Fprintln(c.writer, "Error removing Book! :$v", err.Error())
	} else {
		fmt.Fprintln(c.writer, "Book removed successfully!")
	}
}

func (c *Console) borrowBook(scanner *bufio.Scanner) {
	fmt.Fprint(c.writer, "Enter book ID to borrow: ")
	scanner.Scan()
	bookIDStr := strings.TrimSpace(scanner.Text())
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid book ID. Please enter a number.")
		return
	}

	fmt.Fprint(c.writer, "Enter member ID: ")
	scanner.Scan()
	memberIDStr := strings.TrimSpace(scanner.Text())
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid member ID. Please enter a number.")
		return
	}

	err = c.libraryService.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Fprintln(c.writer, "Error borrowing book:", err)
		return
	}

	fmt.Fprintln(c.writer, "Book borrowed successfully!")
}

func (c *Console) returnBook(scanner *bufio.Scanner) {
	fmt.Fprint(c.writer, "Enter book ID to return: ")
	scanner.Scan()
	bookIDStr := strings.TrimSpace(scanner.Text())
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid book ID. Please enter a number.")
		return
	}

	fmt.Fprint(c.writer, "Enter member ID: ")
	scanner.Scan()
	memberIDStr := strings.TrimSpace(scanner.Text())
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid member ID. Please enter a number.")
		return
	}

	err = c.libraryService.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Fprintln(c.writer, "Error returning book:", err)
		return
	}

	fmt.Fprintln(c.writer, "Book returned successfully!")
}

func (c *Console) listAvailableBooks() {
	books := c.libraryService.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Fprintln(c.writer, "No available books.")
		return
	}

	fmt.Fprintln(c.writer, "Available books:")
	for _, book := range books {
		fmt.Fprintf(c.writer, "ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID(), book.Title(), book.Author(), book.Status())
	}
}

func (c *Console) listBorrowedBooks(scanner *bufio.Scanner) {
	fmt.Fprint(c.writer, "Enter member ID: ")
	scanner.Scan()
	memberIDStr := strings.TrimSpace(scanner.Text())
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Fprintln(c.writer, "Invalid member ID. Please enter a number.")
		return
	}

	books := c.libraryService.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Fprintln(c.writer, "No borrowed books.")
		return
	}

	fmt.Fprintln(c.writer, "Borrowed books:")
	for _, book := range books {
		fmt.Fprintf(c.writer, "ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID(), book.Title(), book.Author(), book.Status())
	}
}

