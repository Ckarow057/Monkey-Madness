module monkey_madness/main_api

require (
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.22
	monkey_madness/monkey_database v0.0.0-00010101000000-000000000000
)

require monkey_madness/monkey_models v0.0.0-00010101000000-000000000000 // indirect

go 1.22.2

replace monkey_madness/monkey_database => ../database

replace monkey_madness/monkey_models => ../models
