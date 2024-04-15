package monkeydb

import (
	"database/sql"
	"monkey_madness/monkey_models"
	"log"
)

func getAllMonkeys(db *sql.DB) []models.Monkey