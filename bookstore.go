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

func InitializeDatabase(dbPath string) (*sql.DB, error) {

}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
