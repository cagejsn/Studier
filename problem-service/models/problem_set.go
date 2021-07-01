package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/beevik/guid"
)

type ProblemSet struct {
	ID               string          `json:"id"`
	AccessLevel      string          `json:"accessLevel"`
	Title            string          `json:"title"`
	IncludedProblems json.RawMessage `json:"includedProblems,omitempty"`
	Author           string          `json:"author"`
}

func CreateProblemSet(problemSet *ProblemSet) (*ProblemSet, error) {

	guidString := guid.NewString()
	_, err := db.Exec("insert into problem_set (id) values ($1)", guidString)
	if err != nil {
		return nil, err
	}

	probSet := ProblemSet{
		ID: guidString,
	}

	return &probSet, nil

}

func GetProblemSet(id string) (*ProblemSet, error) {
	row := db.QueryRow("SELECT id, access_level, title, included_problems, author FROM problem_set where id = $1", id)

	probSet := ProblemSet{
		// IncludedProblems: (*json.RawMessage)(new([]byte)),
	}

	err := row.Scan(&probSet.ID, &probSet.AccessLevel, &probSet.Title, (*[]byte)(&probSet.IncludedProblems), &probSet.Author)
	if err != nil {
		return nil, err
	}

	return &probSet, nil
}

func UpdateProblemSet(problemSet *ProblemSet) error {

	sqlStatement := `UPDATE problem_set SET title = ?, access_level = ?, included_problems = ?, author = ? WHERE id = ?`
	result, err := db.Exec(sqlStatement, problemSet.Title, problemSet.AccessLevel, problemSet.IncludedProblems, problemSet.Author, problemSet.ID)
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
