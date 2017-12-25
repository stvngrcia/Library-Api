package main

import (
	"database/sql"
	"errors"
)
/*Defines the books*/
type books struct {
	ID    int   `json:"id"`
	NAME  string `json:"name"`
	DESCRIPTION string `json:"description"`
}

/*Gets a book in the database*/
func (b *books) getBooks(db *sql.DB) error{
	return errors.New("Not yet implemented.")
}

/*Updates a book info*/
func (b *books) updateBooks(db *sql.DB) error{
	return errors.New("Not yet implemented.")
}

/*Deletes Books*/
func (b *books) deleleteBooks(db *sql.DB) error{
	return errors.New("No yet implemented.")
}

/*Adds books*/
func (b *books) addBooks(db *sql.DB) error{
	return errors.New("No yet implemented.")
}

/*Gets all the boooks in the database*/
func (b *books) getAllBooks(db *sql.DB, start, count int)([]books, error){
	return nil, errors.New("No yet implemented.")
}

