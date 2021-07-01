package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/cagejsn/Studier/problem-service/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func registerProblemSetHandlers(r *mux.Router) {

	r.HandleFunc("/problem-set", func(w http.ResponseWriter, r *http.Request) {
		problemSet, err := models.CreateProblemSet(&models.ProblemSet{})
		log.Print(err)
		v, _ := json.Marshal(problemSet)
		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("POST")

	r.HandleFunc("/problem-set/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		problem, err := models.GetProblemSet(id)
		if err != nil {
			log.Panic(err)
		}

		log.Println(problem)
		v, err := json.Marshal(problem)
		if err != nil {
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("GET")

	r.HandleFunc("/problem-set", func(w http.ResponseWriter, r *http.Request) {

		var pS models.ProblemSet
		err := json.NewDecoder(r.Body).Decode(&pS)
		if err != nil {
			log.Fatal("bad request", err)
			io.WriteString(w, "BAD REQUEST")
		}

		err = models.UpdateProblemSet(&pS)
		if err != nil {
			log.Panic(err)
		}

		updatedProblemSet, err := models.GetProblemSet(pS.ID)
		if err != nil {
			log.Panic(err)
		}

		v, err := json.Marshal(updatedProblemSet)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))

	}).Methods("PUT")

	// r.HandleFunc("/problem-set/{id}", func(w http.ResponseWriter, r *http.Request) {

	// 	vars := mux.Vars(r)
	// 	id := vars["id"]

	// 	var pS models.ProblemSet
	// 	err := json.NewDecoder(r.Body).Decode(&pS)
	// 	if err != nil {
	// 		log.Fatal("bad request", err)
	// 		io.WriteString(w, "BAD REQUEST")
	// 	}

	// 	// updatedProblemSet, err := models.UpdateProblemSet(&pS)
	// 	// if err != nil {
	// 	// 	log.Panic(err)
	// 	// }

	// 	v, err := json.Marshal(updatedProblemSet)
	// 	if err != nil {
	// 		io.WriteString(w, "INTERNAL SERVER ERROR")
	// 		log.Panic(err)
	// 	}

	// 	enableCors(&w)
	// 	io.WriteString(w, string(v))

	// }).Methods("PATCH")
}
