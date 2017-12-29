package main
import (
	"encoding/json"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/**
 * getBook - Gets a single book based on its id.
 * Description: Maps the request to get the mux variable id to then call
 * a book method to query the db.
 */
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

/**
 * getAllBooks - Gets books in the data base from start up to count. By default
 * it will start at 0. No more the 10 items will be displayed per query.
 * Return: Upon sucess a map with the books and nil as error. Otherwise a nil
 * map and the error message.
 */
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
