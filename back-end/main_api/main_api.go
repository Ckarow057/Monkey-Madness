package main

import (
	"log"
	"net/http"

	"database/sql"
	// "encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var monkeyDatabase *sql.DB

func main() {
	router := mux.NewRouter()

	var err error
	monkeyDatabase, err = sql.Open("sqlite3", "../database/monkey.db")

	if err != nil {
		log.Fatal(err)
	}
	defer monkeyDatabase.Close()

	router.HandleFunc("/", hello).Methods(http.MethodPost)
	// router.HandleFunc("/dailymonkey", hello).Methods(http.MethodPost)
	// router.HandleFunc("/displayallmonkeys", hello).Methods(http.MethodPost)
	// router.HandleFunc("/displaymonkey", hello).Methods(http.MethodPost)
	// router.HandleFunc("/displaypersonalinfo", hello).Methods(http.MethodPost)
	// router.HandleFunc("/insertuserinfo", hello).Methods(http.MethodPost)


	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Monkey Madness Time!"))

}

