# Library Management System

## Introduction

The Library Management System is a command-line application written in Go that allows users to manage books and members in a library. It supports operations like adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

## Features

- Add a new book to the library
- Remove an existing book from the library
- Borrow a book for a member
- Return a borrowed book
- List all available books
- List all borrowed books by a member

## Installation

To install and run the Library Management System, follow these steps:

1. **Clone the repository**:

   ```sh
   git clone https://github.com/beka-birhanu/library_management.git
   cd library_management
   ```

2. **Build the application**:

   ```sh
   make build
   ```

3. **Run the application**:
   ```sh
   make run
   ```

## Usage

The Library Management System provides a console-based interface. Upon running the application, you will see a menu with various options:

1. Add a new book
2. Remove an existing book
3. Borrow a book
4. Return a book
5. List all available books
6. List all borrowed books by a member
7. Exit

### Add a New Book

To add a new book, select option `1` and provide the required details (ID, title, author, and status).

### Remove an Existing Book

To remove a book, select option `2` and provide the book ID.

### Borrow a Book

To borrow a book, select option `3`, provide the book ID, and the member ID.

### Return a Book

To return a book, select option `4`, provide the book ID, and the member ID.

### List All Available Books

To list all available books, select option `5`.

### List All Borrowed Books by a Member

To list all borrowed books by a member, select option `6` and provide the member ID.

### Exit

To exit the application, select option `7`.

## Project Structure

The project is organized as follows:

- `models`: Contains the data models for `Book` and `Member`.
- `common`: Defines the `LibraryManager` interface.
- `services`: Contains the `Library` implementation that manages the library's operations.
- `controllers`: Contains the `Console` controller for the console-based user interface.
- `main.go`: Entry point of the application.
