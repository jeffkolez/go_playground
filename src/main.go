package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./response"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// homeHandler GET handler for home
func homeHandler(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]string)
	ret["name"] = "jeff"
	response.RespondWithJSON(w, http.StatusOK, ret)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB")
	log.Print(host)

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).
		Methods("GET")

	var addr string
	addr = ":" + os.Getenv("APP_PORT")

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
