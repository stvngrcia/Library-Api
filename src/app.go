package main

import (
	"fmt"
	"log"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_"github.com/lib/pq"
)

/**
 * App - Structure that Exposes references to the router and db.
 */
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

/**
 * Initialize - Creates db connection and initializes router/
 * @user: Db user name.
 * @password: Db password.
 * @dbname: Name of the db.
 * Upon error the log.Fatal will write message to stderr and exit with 1.
 */
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

/**
 * Run - starts the application at the given port.
 * @port: String representing the port in which the app will be running.
 */
func (app *App) Run(port string){
	log.Fatal(http.ListenAndServe(port, app.Router))
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
