package main

import (
	"log"
	"net/http"

	"database/sql"
	"monkey_madness/monkey_database"

	"encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var monkeyDatabase *sql.DB

func main() {
	router := mux.NewRouter()

	var err error
	monkeyDatabase, err = sql.Open("sqlite3", "../database/monkeys.db")

	if err != nil {
		log.Fatal(err)
	}
	defer monkeyDatabase.Close()

	router.HandleFunc("/", hello).Methods(http.MethodPost)
	// router.HandleFunc("/dailymonkey", hello).Methods(http.MethodPost)
	router.HandleFunc("/displayallmonkeys", getMonkeysHandler).Methods(http.MethodPost)
	// router.HandleFunc("/displaymonkey", hello).Methods(http.MethodPost)
	// router.HandleFunc("/displaypersonalinfo", hello).Methods(http.MethodPost)
	// router.HandleFunc("/insertuserinfo", hello).Methods(http.MethodPost)


	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Monkey Madness Time!"))

}

func getMonkeysHandler(w http.ResponseWriter, r *http.Request){
	monkeys := monkey_database.GetAllMonkeys(monkeyDatabase)

	monkeyJSON, err := json.Marshal(monkeys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)
	w.Write(monkeyJSON)
}

func enableCORS(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}