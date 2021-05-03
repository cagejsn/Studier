package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/beevik/guid"
)

type Grade struct {
	ID               string           `json:"gradeId"`
	Version          int              `json:"version"`
	ProblemAttemptID string           `json:"problemAttemptId"`
	GradeType        string           `json:"gradeType"`
	Outcome          *json.RawMessage `json:"outcome"`
	SubmissionState  string           `json:"submissionState"`
}

func SaveGrade(grade *Grade) (*Grade, error) {

	//validate the problemAttempt
	if grade.ProblemAttemptID == "" {
		var err = errors.New("no problemAttemptID given for grade")
		return nil, err
	}

	guidString := guid.NewString()
	_, err := db.Exec("insert into grade (id, problem_attempt_id) values ($1, $2)", guidString, grade.ProblemAttemptID)
	if err != nil {
		return nil, err
	}

	return GetGrade(guidString)
}

func GetGrade(id string) (*Grade, error) {
	row := db.QueryRow("SELECT id, version, problem_attempt_id, grade_type, outcome, submission_state FROM grade where id = $1", id)

	g := Grade{
		Outcome: (*json.RawMessage)(new([]byte)),
	}

	err := row.Scan(&g.ID, &g.Version, &g.ProblemAttemptID, &g.GradeType, (*[]byte)(g.Outcome), &g.SubmissionState)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func UpdateGrade(grade *Grade) (*Grade, error) {

	//TODO: write better validations that actually take into account the various falsy go values
	//we must know the id
	if grade.ID == "" {

		var err = errors.New("which id to update?")
		return nil, err
	}

	// TODO: because of the sql that we used, we need both outcome and submissionState, we can make orthagonal updates at some point
	//either the userInput or the SubmissionState should be present, or else update is pointless
	if grade.Outcome == nil || grade.SubmissionState == "" {

		var err = errors.New("there is a issue with the outcome or submission state")
		return nil, err
	}

	sqlStatement := `UPDATE grade SET outcome = ?, submission_state = ? WHERE id = ?`
	result, err := db.Exec(sqlStatement, grade.Outcome, grade.SubmissionState, grade.ID)
	if err != nil {
		return nil, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(lastID)
	return GetGrade(grade.ID)
}
