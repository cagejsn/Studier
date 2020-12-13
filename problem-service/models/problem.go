package models

import "github.com/beevik/guid"

type Problem struct {
	ID        string `json:"id"`
	Version   int    `json:"version"`
	Statement string `json:"statement"`
	Type      string `json:"type"`
	Answers   string `json:"answers"`
}

func GetAllProblems() ([]*Problem, error) {
	rows, err := db.Query("SELECT * FROM problem")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	probs := make([]*Problem, 0)
	for rows.Next() {
		prob := new(Problem)
		err := rows.Scan(&prob.ID, &prob.Version, &prob.Statement, &prob.Type, &prob.Answers)
		if err != nil {
			return nil, err
		}
		probs = append(probs, prob)
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

	prob := Problem{}

	err := row.Scan(&prob.ID, &prob.Version, &prob.Statement, &prob.Type, &prob.Answers)
	if err != nil {
		return nil, err
	}

	return &prob, nil
}
