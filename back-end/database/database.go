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

func GetRandomMonkey(db *sql.DB) []monkey_models.Monkey{
	numberOfIds, err := countMonkeyIDs(db)
	var monkeyArray[] monkey_models.Monkey
	
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
	monkeyArray = append(monkeyArray, newMonkey)
	return monkeyArray
}

func GetAllUserInfo(db *sql.DB) []monkey_models.PersonalInfo{
	var userInfo[] monkey_models.PersonalInfo

	row, err := db.Query("SELECT * FROM userinformation")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var userId int
		var userSsn int
		var userCardNumber int
		
		row.Scan(&userId, &userSsn, &userCardNumber)
		userData := monkey_models.PersonalInfo {
			UserID: userId,
			UserSSN: userSsn,
			UserCardNum: userCardNumber,
			
		}
		userInfo = append(userInfo, userData)
	}
	return userInfo
}

func InsertUserInfo(db *sql.DB, userdata monkey_models.PersonalInfo) monkey_models.PersonalInfo{
	insertDataSql := "INSERT INTO userinformation (UserID, UserSSN, UserCardNumber) VALUES(?, ?, ?)"
	statement, err := db.Prepare(insertDataSql)

	if err != nil{
		log.Fatal(err)
	}

	result, err := statement.Exec(userdata.UserID, userdata.UserSSN, userdata.UserCardNum)
	if err != nil{
		log.Fatal(err)
	}

	lid, err := result.LastInsertId()
	if err != nil{
		log.Fatal(err)
	}

	userdata.UserID = int(lid)
	return userdata
}