package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/beevik/guid"
)

type ProblemAttempt struct {
	ID              string           `json:"problemAttemptId"`
	Version         int              `json:"version"`
	ProblemID       string           `json:"problemId"`
	UserInput       *json.RawMessage `json:"userInput"`
	UserID          string           `json:"userId"`
	SubmissionState string           `json:"submissionState"`
}

func SaveProblemAttempt(problemAttempt *ProblemAttempt) (*ProblemAttempt, error) {

	//validate the problemAttempt
	if problemAttempt.ProblemID == "" {
		var err = errors.New("no problemID given for problemAttempt")
		return nil, err
	}

	guidString := guid.NewString()
	_, err := db.Exec("insert into problem_attempt (id, problem_id, person) values ($1, $2, 'cage')", guidString, problemAttempt.ProblemID)
	if err != nil {
		return nil, err
	}

	// problemAttempt.ID = guidString // TODO: the only action of this method at this point is to add the guid to indicate that the problem-attempt is persisted
	return GetProblemAttempt(guidString)

	// return problemAttempt, nil

}

func GetProblemAttempt(id string) (*ProblemAttempt, error) {
	row := db.QueryRow("SELECT id, version, problem_id, user_input, submission_state FROM problem_attempt where id = $1", id)

	pA := ProblemAttempt{
		UserInput: (*json.RawMessage)(new([]byte)),
	}

	err := row.Scan(&pA.ID, &pA.Version, &pA.ProblemID, (*[]byte)(pA.UserInput), &pA.SubmissionState)
	if err != nil {
		return nil, err
	}

	return &pA, nil
}

func UpdateProblemAttempt(problemAttempt *ProblemAttempt) (*ProblemAttempt, error) {

	//TODO: write better validations that actually take into account the various falsy go values
	//we must know the id
	if problemAttempt.ID == "" {

		var err = errors.New("which id to update?")
		return nil, err
	}

	// TODO: because of the sql that we used, we need both userInput and submissionState, we can make orthagonal updates at some point
	//either the userInput or the SubmissionState should be present, or else update is pointless
	if problemAttempt.UserInput == nil || problemAttempt.SubmissionState == "" {

		var err = errors.New("there is a issue with the user input or submission state")
		return nil, err
	}

	sqlStatement := `UPDATE problem_attempt SET user_input = ?, submission_state = ? WHERE id = ?`
	result, err := db.Exec(sqlStatement, problemAttempt.UserInput, problemAttempt.SubmissionState, problemAttempt.ID)
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
	return GetProblemAttempt(problemAttempt.ID)
}
