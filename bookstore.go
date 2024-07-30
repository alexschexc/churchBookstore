package bookstore

import (
	"errors"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Book struct {
	ISBN       string // database Book PK, ISBN-13
	Title      string
	Author     string
	Copies     int
	PubVersion string
	PubYear    string
}
type PrayerRope struct {
	Steps        int
	Subdivisions string
	Quantity     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}
