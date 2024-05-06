package main

import (
	"log"
	"net/http"

	"database/sql"
	"monkey_madness/monkey_database"
	"monkey_madness/monkey_models"

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

	router.HandleFunc("/displayallmonkeys", getMonkeysHandler).Methods(http.MethodPost) // Displays every monkey in db
	router.HandleFunc("/displaymonkey", getMonkeyHandler).Methods(http.MethodPost) // displays random monkey
	router.HandleFunc("/displaymonkeybyname", getMonkeyByNameHandler).Methods(http.MethodPost) // displays specifiic monkey by name

	router.HandleFunc("/displaypersonalinfo", getAllUserInfoHandler).Methods(http.MethodPost)
	router.HandleFunc("/insertuserinfo", addUserInfoHandler).Methods(http.MethodPost)

	http.ListenAndServe(":8000", router)
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

func getMonkeyHandler(w http.ResponseWriter, r *http.Request){
	monkeys := monkey_database.GetRandomMonkey(monkeyDatabase)

	monkeyJSON, err := json.Marshal(monkeys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)
	w.Write(monkeyJSON)
}

func getMonkeyByNameHandler(w http.ResponseWriter, r *http.Request){
	var monkey monkey_models.Monkey

	err:= json.NewDecoder(r.Body).Decode(&monkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	newMonkey := monkey_database.GetMonkeyByName(monkeyDatabase, monkey)

	newMonkeyJSON, err := json.Marshal(newMonkey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)
	w.Write(newMonkeyJSON)
}


func getAllUserInfoHandler(w http.ResponseWriter, r *http.Request){
	userData := monkey_database.GetAllUserInfo(monkeyDatabase)

	userJSON, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)
	w.Write(userJSON)
}

func addUserInfoHandler(w http.ResponseWriter, r *http.Request){
	var userdata monkey_models.PersonalInfo

	err:= json.NewDecoder(r.Body).Decode(&userdata)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Put error handling here to check for missing params (Bad JSON)

	newUserData := monkey_database.InsertUserInfo(monkeyDatabase, userdata)

	newDataJSON, err := json.Marshal(newUserData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enableCORS(&w)
	w.Write(newDataJSON)
}
func enableCORS(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

