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
	})

	r.HandleFunc("/problem/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		enableCors(&w)
		problem, err := models.GetProblem(id)
		if err != nil {
			log.Panic(err)
		}

		v, _ := json.Marshal(problem)
		io.WriteString(w, string(v))
	})

	startServerTLS(r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")

	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization")

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
