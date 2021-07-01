package main

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type SampleContentType struct {
	SampleContentType string
	SampleContent     *json.RawMessage
}

func getSampleContentWithTypeForUser(sC string, userID string) (*SampleContentType, error) {
	return &SampleContentType{
		sC,
		nil,
	}, nil
}
