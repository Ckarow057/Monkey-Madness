package monkey_database

import (
	"database/sql"
	"monkey_madness/monkey_models"
	"log"
	"math/rand"
)

func GetAllMonkeys(db *sql.DB) []monkey_models.Monkey{
	var monkeys[] monkey_models.Monkey

	row, err := db.Query("SELECT * FROM monkeys")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var monkeyId int
		var monkeyName string
		var monkeyBreed string
		var monkeyImg string
		var monkeyFact string
		row.Scan(&monkeyId, &monkeyName, &monkeyBreed, &monkeyImg, &monkeyFact)
		monkey := monkey_models.Monkey {
			MonkeyID: monkeyId,
			MonkeyName: monkeyName,
			MonkeyBreed: monkeyBreed,
			MonkeyImg: monkeyImg,
			MonkeyFact: monkeyFact,
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func GetMonkeyByName(db *sql.DB, monkey monkey_models.Monkey) monkey_models.Monkey{
	sqlStatement := `SELECT * FROM monkeys WHERE monkeyName = ?`
	row := db.QueryRow(sqlStatement, monkey.MonkeyName)
	var monkeyId int
	var monkeyName string
	var monkeyBreed string
	var monkeyImg string
	var monkeyFact string
	row.Scan(&monkeyId, &monkeyName, &monkeyBreed, &monkeyImg, &monkeyFact)
	newMonkey := monkey_models.Monkey {
		MonkeyID: monkeyId,
		MonkeyName: monkeyName,
		MonkeyBreed: monkeyBreed,
		MonkeyImg: monkeyImg,
		MonkeyFact: monkeyFact,
	}
	return newMonkey
}

func countMonkeyIDs(db *sql.DB) (int, error) {
	var count int
	err:= db.QueryRow("SELECT COUNT(*) FROM monkeys").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count, err
}

func GetRandomMonkey(db *sql.DB) monkey_models.Monkey{
	numberOfIds, err := countMonkeyIDs(db)
	
	if err != nil {
		log.Fatal(err)
	}
	
	randomIndex:= rand.Intn(numberOfIds) + 1

	sqlStatement := `SELECT * FROM monkeys WHERE monkeyID = ?`
	row := db.QueryRow(sqlStatement, randomIndex)
	var monkeyId int
	var monkeyName string
	var monkeyBreed string
	var monkeyImg string
	var monkeyFact string
	row.Scan(&monkeyId, &monkeyName, &monkeyBreed, &monkeyImg, &monkeyFact)
	newMonkey := monkey_models.Monkey {
		MonkeyID: monkeyId,
		MonkeyName: monkeyName,
		MonkeyBreed: monkeyBreed,
		MonkeyImg: monkeyImg,
		MonkeyFact: monkeyFact,
	}
	return newMonkey
}