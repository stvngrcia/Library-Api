package main

import "os"

/*Entry point. Gets envs for database and runs application on a port*/
func main(){
	app := App{}
	app.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"))
	app.Run(":80")
}