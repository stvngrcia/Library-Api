package main

import (

	"encoding/json"
	"fmt"
	"log"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_"github.com/lib/pq"
)
/* Exposes references to the router and the database for the App.*/
type App struct {
	Router *mux.Router
	DB     *sql.DB
}


/*Takes in the methods to initilize the database and connect to it.*/
func (app *App) Initialize(user, password, dbname string){
	connect :=
		fmt.Sprintf("user=%s password=%s dbname=%s",user, password, dbname)
	var err error
	app.DB, err = sql.Open("postgres", connect)
	if err != nil{
		log.Fatal(err)
	}
	app.Router = mux.NewRouter().StrictSlash(true)
	app.initializeRoutes()
}

/* starts the application at the given port.*/
func (app *App) Run(port string){
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

/**
 * initializeRoutes - Initilizes the endpoints for each specific request.
 */
func (app *App) initializeRoutes(){
	app.Router.HandleFunc("/books", app.getBooks).Methods("GET")
	app.Router.HandleFunc("/book", app.createBook).Methods("POST")
	app.Router.HandleFunc("/book/{id:[0-9]+}", app.getBook).Methods("GET")
	app.Router.HandleFunc("/book/{id:[0-9]+}", app.updateBook).Methods("PUT")
	app.Router.HandleFunc("/book/{id:[0-9]+}", app.deleteBook).Methods("DELETE")
}

/*Handler to fetch a single book*/
func (app *App) getBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	b := books{ID: id}
	if err := b.getBook(app.DB); err != nil{
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Book not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, b)
}

/*Gets first 10 books*/
func (app *App) getBooks(w http.ResponseWriter, r *http.Request){
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1{
		count = 10
	}
	if start < 0{
		start = 0
	}
	book, err := getAllBooks(app.DB, start, count)
	if err != nil{
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	respondWithJSON(w, http.StatusOK, book)
}

/**
 * createBook - Adds a book to the database.
 * @w: The response writter
 * @r: The response
 * Description: Assumes that the response body is a JSON object containing the
 * details of the product created. It extracts the object into a book and uses
 * the addBooks method to create a book.
 */
func (app *App) createBook(w http.ResponseWriter, r *http.Request){
	var b books
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid Request payload")
		return
	}
	fmt.Println(b)
	defer r.Body.Close()
	if err := b.addBooks(app.DB); err != nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, b)
}

/**
 * UpdateBook - Updates an existing book in the database
 * @w: The response writter
 * @r: The response
 * Description: Extracts the product details from the request body as well as
 * the id from the URL. Uses this id to update the book.
 */
func (app *App) updateBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		respondWithError(w, http.StatusNotFound, "Invalid id")
		return
	}
	var b books
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil{
		respondWithError(w, http.StatusBadRequest, "invalid payload")
		return
	}
	defer r.Body.Close()
	b.ID = id

	if err := b.updateBooks(app.DB); err != nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, b)
}

/**
 * delteBook - Removes a book from the database.
 * @w: The response writter
 * @r: The response
 * Description: Extracts the id from the resquest URL and uses it to delete
 * the corresponding product from the database.
 *
 */
func (app *App) deleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "Invalid Book id")
		return
	}

	b := books{ID: id}
	if err := b.deleteBooks(app.DB); err != nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

/*Handling the response for errors*/
func respondWithError(w http.ResponseWriter, code int, message string){
	respondWithJSON(w, code, map[string]string{"error": message})
}

/*Handles conversion to json*/
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}