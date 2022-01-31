package models

import "errors"

// ErrNoRecord ...
var ErrNoRecord = errors.New("models: подходящей записи не найдено")

// Snippet ...
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created string
	Expires string
}
