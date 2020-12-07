package main

import (
	"net/http"
	"io"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kabukky/httpscerts"
	_ "github.com/mattn/go-sqlite3"
	"./models"
	"encoding/json"
)

func main() {

	models.OpenDatabaseConnection("./problems.db")

	http.HandleFunc("/content", func(w http.ResponseWriter,r *http.Request){
		enableCors(&w)
		probs, _ := models.AllProblems()
		v, _ := json.Marshal(probs)
		io.WriteString(w, string(v))
	})
	startServerTLS()
}


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")

	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization")

}


func startServerTLS(){
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8090")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	http.ListenAndServeTLS(":8090", "cert.pem", "key.pem", nil)
}


