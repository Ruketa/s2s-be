package db

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type questionnaire struct{
	Id string `json:"id"`
	Satisfaction_level int `json:"satisfaction_level"`
	Recommendation_level int `json:"recommendation_level"`
	Topics string `json:"topics"`
	Participation_level int `json:"participation_level"`
	Presentation_level int `json:"presentation_level"`
	Free_comment string `json:"free_comment"`
	Holding_num int `json:"holding_num"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Connect() *sql.DB {
	const (
		HOST		 = "*"
		DATABASE = "*"
		USER     = "*"
		PORT 		 = 8080 
		PASSWORD = "*"
	)

	// initialize connection string
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)
	fmt.Println(connectionString)

	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	return db
}

func CheckConnection(db *sql.DB) {
	var err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")
}

func Close(db *sql.DB) {
	db.Close()
}

func CreateTable(db *sql.DB) {
	var err = db.Ping()
	checkError(err)

	// questionnaire table
	var query string = "CREATE TABLE IF NOT EXISTS questions ( ";
	query = query + "id SERIAL PRIMARY KEY, satisfaction_level INT NOT NULL, ";
	query = query + "recommendation_level INT NOT NULL, "
	query = query + "topics Text, "
	query = query + "participation_level INT NOT NULL, "
	query = query + "presentation_level INT NOT NULL, "
	query = query + "free_comment Text, "
	query = query + "holding_num INT NOT NULL,"
	query = query + ")"
			
	_, err = db.Exec(query)
	checkError(err)

	// presenter table

}

func SelectQuesionnaire(db *sql.DB)  []questionnaire {
	sql_select := "SELECT * from questionnaire;"
	rows, err := db.Query(sql_select)
	checkError(err)
	defer rows.Close()

	var qs []questionnaire;

	for rows.Next() {
		var q questionnaire;
		err := rows.Scan(
			&q.Id, 
			&q.Satisfaction_level,
			&q.Recommendation_level, 
			&q.Topics, 
			&q.Participation_level, 
			&q.Presentation_level, 
			&q.Free_comment, 
			&q.Holding_num);
		checkError(err)

		qs = append(qs, q)
	}

	return qs
}

func SelectQuesionnaireByHoldingNum(db *sql.DB, holding_num string)  []questionnaire {
	sql_select := "SELECT * from questionnaire WHERE holding_num = " + string(holding_num);
	rows, err := db.Query(sql_select)
	checkError(err)
	defer rows.Close()

	var qs []questionnaire;

	for rows.Next() {
		var q questionnaire;
		err := rows.Scan(
			&q.Id, 
			&q.Satisfaction_level,
			&q.Recommendation_level, 
			&q.Topics, 
			&q.Participation_level, 
			&q.Presentation_level, 
			&q.Free_comment, 
			&q.Holding_num);
		checkError(err)

		qs = append(qs, q)
	}

	return qs
}