package models

type Problem struct {
	Id string
    Statement  string
    Type  string
    Answers string
}

func AllProblems() ([]*Problem, error) {
    rows, err := db.Query("SELECT * FROM problem")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    probs := make([]*Problem, 0)
    for rows.Next() {
        prob := new(Problem)
        err := rows.Scan(&prob.Id, &prob.Statement, &prob.Type, &prob.Answers)
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