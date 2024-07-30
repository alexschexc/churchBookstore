package bookstore

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gdamore/tcell/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rivo/tview"
)

type Book struct {
	ISBN       string // database Book PK, ISBN-13
	Title      string
	Author     string
	Copies     int
	PubVersion string
	PubYear    string
	Price      float64
}
type PrayerRope struct {
	Steps        int
	Material     string
	Subdivisions string
	Quantity     int
}

/*
 other product Structs to added later
*/
// InitializeDatabase sets up the database connection and ensures the schema is correct.
func InitializeDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Ensure the books table has the correct schema
	err = ensureBookSchema(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ensureSchema checks if the table exists and has the necessary columns.
func ensureBookSchema(db *sql.DB) error {
	// Create table if it doesn't exist
	const createTableQuery string = `CREATE TABLE IF NOT EXISTS books (isbn TEXT PRIMARY KEY, title TEXT NOT NULL, author TEXT NOT NULL, copies INTEGER NOT NULL, pubversion TEXT NOT NULL, pubyear TEXT NOT NULL, price REAL);`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

// Buy decrements the copies of the books and returns the updated book
func BuyBook(db *sql.DB, isbn string) (Book, error) {
	var book Book

	// Retrieve the book from the DB
	query := "SELECT copies FROM books WHERE isbn = ?"
	err := db.QueryRow(query, isbn).Scan(&book.Copies)
	if err != nil {
		if err == sql.ErrNoRows {
			return Book{}, errors.New("book not found")
		}
		return Book{}, err
	}
	// Check if there are any copies left to buy
	if book.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	// Decrement the copies count
	book.Copies--

	// Update the book's copies in the database
	updateQuery := "UPDATE books SET copies = ? WHERE isbn = ?"
	_, err = db.Exec(updateQuery, book.Copies, isbn)
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

// Example function to update the price of a book
func UpdatePrice(db *sql.DB, isbn string, price float64) error {
	_, err := db.Exec("UPDATE books SET price = ? WHERE isbn = ?", price, isbn)
	return err
}

// Example function to add a new book to the database
func AddBook(db *sql.DB, book Book) error {
	query := "INSERT INTO books (isbn, title, author, copies, pubversion, pubyear, price) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, book.ISBN, book.Title, book.Author, book.Copies, book.PubVersion, book.PubYear, book.Price)
	return err
}

// Example function to get a book by ISBN
func GetBook(db *sql.DB, isbn string) (Book, error) {
	var book Book
	query := "SELECT isbn, title, author, copies, pubversion, pubyear, price FROM books WHERE isbn = ?"
	err := db.QueryRow(query, isbn).Scan(&book.ISBN, &book.Title, &book.Author, &book.Copies, &book.PubVersion, &book.PubYear, &book.Price)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}
