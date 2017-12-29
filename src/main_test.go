package main_test

import (
	"os"
	"testing"
	"log"
	"."
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"strconv"
)

var app main.App
/**
curl -i -X POST -H 'Content-Type: application/json' -d '{"name":"test", "description":"test1"}' http://localhost:8080/book
**/
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


/*Deletes the records in the db and then checks
  for the respose of a GET request*/
func TestEmptyTable(t *testing.T){
	clearTable()

	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]"{
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

/*Executes the request using the app router and returns the response*/
func executeRequest(req *http.Request) *httptest.ResponseRecorder{
	resp := httptest.NewRecorder()
	app.Router.ServeHTTP(resp, req)
	return resp
}

/*Compares the expected status code against the obtained one.*/
func checkResponseCode(t *testing.T, expected, actual int){
	if expected != actual{
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}


/*Check for a non existant book. Should get 404 and a Book not found message*/
func TestGetNonExistentBook(t *testing.T){
	clearTable()

	req, _ := http.NewRequest("GET", "/book/200", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Book not found"{
		t.Errorf("Expected the 'error key of the response to be set to'Book not found'. Got %s", m["error"])
	}
}

/*Test for the addition of a book*/
func TestCreateBook(t *testing.T){
	clearTable()

	payload := []byte(`{"name":"test book", "description":"Testing this book"}`)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test book"{
		t.Errorf("Expected test book, Got %s instead", m["name"])
	}

	if m["description"] != "Testing this book"{
		t.Errorf("Expected Testing this book, Got %s instead", m["description"])
	}
}

/*Testing for fetching a book*/
func TestGetBook(t *testing.T){
	clearTable()
	addBooks(1)
	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

/*Adding dummy books to the database*/
func addBooks(number int){
	if number < 1{
		number = 1
	}
	for i := 0; i < number; i++{
		app.DB.Exec("INSERT INTO books(name, description) VALUES($1 %2)",
					"Book"+strconv.Itoa(i), (i+1.0)*10)
	}
}

/*Updating a book*/
func TestUpdateBook(t *testing.T){
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := executeRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	payload := []byte(`{"name":"test updated name","description":"updated"}`)
	req, _ = http.NewRequest("PUT", "/book/1", bytes.NewBuffer(payload))
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] == originalBook["name"]{
		t.Errorf("Expected the name to change, but it did not")
	}
}

/*Deleting a book and testing if it was deleted*/
func TestDeleteBook(t *testing.T){
	clearTable()
	addBooks(1)

	req, _ := http.NewRequest("DELETE", "/book/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/book/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}