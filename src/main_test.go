package main_test

import (
	"os"
	"testing"
	"log"
	"."
)

var app main.App

/*Initializing the app to be tested*/
func TestMain(m *testing.M){
	app = main.App{}
	app.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"))
	ensureTableExist()

	code := m.Run()
	clearTable()
	os.Exit(code)
}

/*Checks that the database table exists*/
func ensureTableExist(){
	if _, err := app.DB.Exec(tableCreationQuery); err != nil{
		log.Fatal(err)
	}
}

/*Deletes a table after tests have run*/
func clearTable(){
	app.DB.Exec("DELETE FROM products")
	app.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}

/*Query to create a table inside test db*/
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books(
	id SERIAL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
	)`