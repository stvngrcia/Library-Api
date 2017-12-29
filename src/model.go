package main

import "database/sql"

/**
 * books - Defines the fields for db.
 */
type books struct {
	ID    int   `json:"id"`
	Name  string `json:"name"`
	Description string `json:"description"`
}

/**
 * getBook - books method that quearies the db
 * for a book that contains a specific id.
 */
func (b *books) getBook(db *sql.DB) error{
	return db.QueryRow("SElECT name, description FROM books where id=$1",
	b.ID).Scan(&b.Name, &b.Description)
}

/**
 * updateBooks - Updates a book info
 */
func (b *books) updateBooks(db *sql.DB) error{
	_, err := db.Exec("UPDATE books SET name=$1, description=$2 WHERE id=$3",
		b.Name, b.Description, b.ID)
	return err
}

/*Deletes Books*/
func (b *books) deleteBooks(db *sql.DB) error{
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.ID)
	return err
}

/*Adds books*/
func (b *books) addBooks(db *sql.DB) error{
	err := db.QueryRow(
		"INSERT INTO books(name, description) VALUES($1, $2) RETURNING id",
		b.Name, b.Description).Scan(&b.ID)
	if err != nil{
		return err
	}
	return nil
}

/*Gets all the boooks in the database*/
func getAllBooks(db *sql.DB, start, count int)([]books, error){
	rows, err := db.Query(
		"SELECT id, name, description FROM books LIMIT $1 OFFSET $2", count, start)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	book := []books{}
	for rows.Next(){
		var b books
		if err := rows.Scan(&b.ID, &b.Name, &b.Description); err != nil{
			return nil, err
		}
		book = append(book, b)
	}
	return book, nil
}

