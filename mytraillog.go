package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/api/v1/trips", tripsHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
func tripsHandler(w http.ResponseWriter, r *http.Request) {

	type trip struct {
		ID     string  `json:"id"`
		Name   string  `json:"name"`
		Length float64 `json:"length"`
	}
	var trips = []trip{
		{ID: "1", Name: "Bohusleden_1", Length: 23.6},
		{ID: "2", Name: "Kulturstigen", Length: 4.8},
	}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(trips)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
