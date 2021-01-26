package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/beevik/guid"
)

type Problem struct {
	ID        string           `json:"id"`
	State     string           `json:"state"`
	Version   int              `json:"version"`
	Statement string           `json:"statement"`
	Type      *json.RawMessage `json:"type"`
	Answers   *json.RawMessage `json:"answers"`
	Solution  *json.RawMessage `json:"solution"`
}

func GetAllProblems() ([]*Problem, error) {
	rows, err := db.Query("SELECT * FROM problem")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	probs := make([]*Problem, 0)
	for rows.Next() {
		prob := Problem{
			Type:     (*json.RawMessage)(new([]byte)),
			Answers:  (*json.RawMessage)(new([]byte)),
			Solution: (*json.RawMessage)(new([]byte)),
		}
		err := rows.Scan(&prob.ID, &prob.State, &prob.Version, &prob.Statement, (*[]byte)(prob.Type), (*[]byte)(prob.Answers), (*[]byte)(prob.Solution))
		if err != nil {
			return nil, err
		}
		probs = append(probs, &prob)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return probs, nil
}

func CreateProblem(problem *Problem) (*Problem, error) {

	guidString := guid.NewString()
	_, err := db.Exec("insert into problem (id) values ($1)", guidString)
	if err != nil {
		return nil, err
	}

	prob := Problem{
		ID: guidString,
	}

	return &prob, nil

}

func GetProblem(id string) (*Problem, error) {
	row := db.QueryRow("SELECT * FROM problem where id = $1", id)

	prob := Problem{
		Type:     (*json.RawMessage)(new([]byte)),
		Answers:  (*json.RawMessage)(new([]byte)),
		Solution: (*json.RawMessage)(new([]byte)),
	}

	err := row.Scan(&prob.ID, &prob.State, &prob.Version, &prob.Statement, (*[]byte)(prob.Type), (*[]byte)(prob.Answers), (*[]byte)(prob.Solution))
	if err != nil {
		return nil, err
	}

	return &prob, nil
}

func UpdateProblem(problem *Problem) error {

	// tx, _ := db.Begin()
	// stmt, _ := tx.Prepare("UPDATE problem SET state = ?, statement = ? WHERE id = ?")
	// _, err := stmt.Exec(problem.State, problem.Statement, problem.ID)
	// // checkError(err)
	// tx.Commit()

	sqlStatement := `UPDATE problem SET state = ?, statement = ?, type = ?, answers = ?, solution = ? WHERE id = ?`
	result, err := db.Exec(sqlStatement, problem.State, problem.Statement, problem.Type, problem.Answers, problem.Solution, problem.ID)
	if err != nil {
		return err
	}
	log.Println(result.RowsAffected())

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	lastId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(lastId)
	return nil
}
