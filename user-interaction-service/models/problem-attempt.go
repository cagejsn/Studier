package models

import (
	"encoding/json"
	"fmt"

	"github.com/beevik/guid"
)

type ProblemAttempt struct {
	ID                string           `json:"problemAttemptId"`
	ProblemID         string           `json:"problemId"`
	UserInput         *json.RawMessage `json:"userInput"`
	ProblemSolution   *json.RawMessage `json:"problemSolution"`
	SubmissionState   string           `json:"submissionState"`
	SubmissionOutcome *json.RawMessage `json:"submissionOutcome"`
}

func SaveProblemAttempt(problemAttempt *ProblemAttempt) (*ProblemAttempt, error) {

	guidString := guid.NewString()
	_, err := db.Exec("insert into problem_attempt (id, problem_id, person) values ($1, $2, 'cage')", guidString, problemAttempt.ProblemID)
	if err != nil {
		return nil, err
	}

	problemAttempt.ID = guidString // TODO: the only action of this method at this point is to add the guid to indicate that the problem-attempt is persisted

	return problemAttempt, nil

}

func GetProblemAttempt(id string) (*ProblemAttempt, error) {
	row := db.QueryRow("SELECT id, problem_id, user_input, submission_state, submission_outcome FROM problem_attempt where id = $1", id)

	pA := ProblemAttempt{
		UserInput:         (*json.RawMessage)(new([]byte)),
		ProblemSolution:   (*json.RawMessage)(new([]byte)),
		SubmissionOutcome: (*json.RawMessage)(new([]byte)),
	}

	err := row.Scan(&pA.ID, &pA.ProblemID, (*[]byte)(pA.UserInput), &pA.SubmissionState, (*[]byte)(pA.SubmissionOutcome))
	if err != nil {
		return nil, err
	}

	return &pA, nil
}

func UpdateProblemAttempt(problemAttempt *ProblemAttempt) error {

	sqlStatement := `UPDATE problem_attempt SET user_input = ?, submission_state = ?, submission_outcome = ? WHERE id = ?`
	result, err := db.Exec(sqlStatement, problemAttempt.UserInput, problemAttempt.SubmissionState, problemAttempt.SubmissionOutcome, problemAttempt.ID)
	if err != nil {
		return err
	}

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
