package main

import (
	"fmt"
	"log"
	"database/sql"
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
	app.Router = mux.NewRouter()
}

/* starts the application at the given port.*/
func (app *App) Run(port string){}