# NotesRepo

This is the simple forum application written with Go allowing to manage notes repository.


## Setup

####Database
The application uses PostgreSQL database.
In *data/setup.sql* file You will find script to create DB.

####Configuration
The endpoint under which the webservice will be available is set in **config.json** file (Address).

####Golang
To install go visit: https://golang.org/doc/install

## Build 
First You need to install some additional dependencies.
``` 
go get github.com/lib/pq
go get github.com/DATA-DOG/go-sqlmock
go get github.com/stretchr/testify/assert 
```


To finally build the project enter the notesrepo location and type ```go build```
The **notesrepo.exe** executable file will be placed in project location.

## Run 
To run the application just type:
```
./notesrepo
```
In the browser open address set in **config.json** (Address) file (eg. localhost:8080).

