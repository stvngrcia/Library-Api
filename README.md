# Library-Api
Creating a REST API for books at Holberton School, using Go and Postgresql

## Usage
First Install postgres db

```
sudo apt-get install postgresql postgresql-contrib

```

+ Add the following environmental variables:
```

export DB_USERNAME="your_user_name"
export DB_PASSWD="your_passwd"
export DB_NAME="your_db_name"

```
+ to test your application inside the src file simply run:

```
go test -v
```
+ Then if you are running a Linux machine simply go to bin

```
$ ./src
```
Otherwise head over to the src directory and after setting up your GOPATH and your GOBIN run go build or go install.


+ Once the server is running head over to http://localhost:8080/books to see the
books inside your database.

+ If you have not added books yet, you can do manually inside psql or

```
curl -i -X POST -H 'Content-Type: application/json' -d '{"name":"book name", "description":"book description"}' http://localhost:8080/book

```

## TODO

+ add a front end to interact with api