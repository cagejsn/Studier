package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kabukky/httpscerts"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	models.OpenDatabaseConnection("./problems.db")

	r := mux.NewRouter()
	r.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)
		probs, err := models.GetAllProblems()
		log.Print(err)
		v, _ := json.Marshal(probs)
		io.WriteString(w, string(v))
	})

	r.HandleFunc("/problem", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		problem, err := models.CreateProblem(&models.Problem{})
		log.Print(err)
		v, _ := json.Marshal(problem)
		io.WriteString(w, string(v))
	}).Methods("POST")

	r.HandleFunc("/problem/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		enableCors(&w)
		problem, err := models.GetProblem(id)
		if err != nil {
			log.Panic(err)
		}

		log.Println(problem)
		v, err := json.Marshal(problem)
		if err != nil {
			log.Panic(err)
		}

		io.WriteString(w, string(v))
	}).Methods("GET")

	r.HandleFunc("/problem", func(w http.ResponseWriter, r *http.Request) {

		var p models.Problem

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Fatal("bad request", err)
			io.WriteString(w, "BAD REQUEST")
		}

		log.Println(p)
		err = models.UpdateProblem(&p)
		if err != nil {
			log.Panic(err)
		}

		enableCors(&w)
		io.WriteString(w, "OK")
	}).Methods("PUT")

	// the only reason for this options request is that the PUT request seems to do a pre-flight OPTIONS on google chrome
	// and if the CORS doesn't add up on the PREFLIGHT then it will never execute the PUT
	r.HandleFunc("/problem", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		io.WriteString(w, "OK")
	}).Methods("OPTIONS")

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
		err = httpscerts.Generate("cert.pem", "key.pem", "localhost:8090")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	fmt.Print("about to serve")

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "localhost:8090",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())

	http.ListenAndServeTLS(":8090", "cert.pem", "key.pem", r)
}
