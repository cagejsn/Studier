package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/cagejsn/Studier/user-interaction-service/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kabukky/httpscerts"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	models.OpenDatabaseConnection("./userinteraction.db")

	r := mux.NewRouter()

	r.HandleFunc("/problem-attempt", func(w http.ResponseWriter, r *http.Request) {

		var pA models.ProblemAttempt

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&pA)
		if err != nil {
			io.WriteString(w, "BAD REQUEST")
			log.Panic("bad request", err)
		}

		enableCors(&w)
		problemAttempt, err := models.SaveProblemAttempt(&pA)
		if err != nil {
			log.Panic(err)
		}

		v, err := json.Marshal(problemAttempt)
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(w, string(v))
	}).Methods("POST")

	r.HandleFunc("/problem-attempt/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		problemAttempt, err := models.GetProblemAttempt(id)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		v, err := json.Marshal(problemAttempt)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("GET")

	r.HandleFunc("/problem-attempt", func(w http.ResponseWriter, r *http.Request) {

		var pA models.ProblemAttempt

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&pA)
		if err != nil {
			log.Fatal("bad request", err)
			io.WriteString(w, "BAD REQUEST")
		}

		retrievedProblemAttempt, err := models.UpdateProblemAttempt(&pA)
		if err != nil {
			log.Panic(err)
		}

		v, err := json.Marshal(retrievedProblemAttempt)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("PUT")

	// the only reason for this options request is that the PUT request seems to do a pre-flight OPTIONS on google chrome
	// and if the CORS doesn't add up on the PREFLIGHT then it will never execute the PUT
	r.HandleFunc("/problem-attempt", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		io.WriteString(w, "OK")
	}).Methods("OPTIONS")

	r.HandleFunc("/bulk-problem-attempt", func(w http.ResponseWriter, r *http.Request) {

		var pA []models.ProblemAttempt

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&pA)
		if err != nil {
			io.WriteString(w, "BAD REQUEST")
			log.Panic("bad request", err)
		}

		for i, problem := range pA {
			log.Print(problem, i)

		}
		problemAttempt, err := models.SaveProblemAttempt(&models.ProblemAttempt{})
		log.Print(err)
		v, _ := json.Marshal(problemAttempt)

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("POST")

	r.HandleFunc("/grade", func(w http.ResponseWriter, r *http.Request) {

		var g models.Grade

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			io.WriteString(w, "BAD REQUEST")
			log.Panic("bad request", err)
		}

		enableCors(&w)
		grade, err := models.SaveGrade(&g)
		if err != nil {
			log.Panic(err)
		}

		v, err := json.Marshal(grade)
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(w, string(v))
	}).Methods("POST")

	r.HandleFunc("/grade/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		grade, err := models.GetGrade(id)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		v, err := json.Marshal(grade)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("GET")

	r.HandleFunc("/grade", func(w http.ResponseWriter, r *http.Request) {

		//TODO build this with gorilla
		queryParams := r.URL.Query()
		problemAttemptID := queryParams.Get("problemAttemptId")
		log.Print("found problemAttemptId", problemAttemptID)
		if problemAttemptID == "" {
			log.Panic("No problem attempt Id ")
		}

		grade, err := models.GetGrade(problemAttemptID)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		v, err := json.Marshal(grade)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("GET")

	r.HandleFunc("/grade", func(w http.ResponseWriter, r *http.Request) {

		var g models.Grade

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&g)
		if err != nil {
			log.Fatal("bad request", err)
			io.WriteString(w, "BAD REQUEST")
		}

		retrievedGrade, err := models.UpdateGrade(&g)
		if err != nil {
			log.Panic(err)
		}

		v, err := json.Marshal(retrievedGrade)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("PUT")

	// the only reason for this options request is that the PUT request seems to do a pre-flight OPTIONS on google chrome
	// and if the CORS doesn't add up on the PREFLIGHT then it will never execute the PUT
	r.HandleFunc("/grade", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		io.WriteString(w, "OK")
	}).Methods("OPTIONS")

	r.HandleFunc("/sample-content", func(w http.ResponseWriter, r *http.Request) {

		//TODO build this with gorilla
		queryParams := r.URL.Query()
		sampleContentType := queryParams.Get("sampleContentType")
		log.Print("found sampleContentType", sampleContentType)
		if sampleContentType == "" {
			log.Panic("No sampleContentType")
		}

		userID := queryParams.Get("userID")
		log.Print("found userID", userID)
		if userID == "" {
			log.Panic("No userID")
		}

		sC, _ := getSampleContentWithTypeForUser(sampleContentType, userID)
		v, err := json.Marshal(sC)
		if err != nil {
			io.WriteString(w, "INTERNAL SERVER ERROR")
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, string(v))
	}).Methods("GET")

	startServerTLS(r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")

}

func startServerTLS(r *mux.Router) {
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "localhost:8091")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	http.ListenAndServeTLS(":8091", "cert.pem", "key.pem", r)
}
