package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Container struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Tag   string `json:"tag"`
}

var (
	containers = []Container{
		{1, "Web Server", "nginx", "latest"},
		{2, "Database", "postgres", "15-alpine"},
		{3, "Cache", "redis", "7.2"},
		{4, "Application Runtime", "node", "20-alpine"},
		{5, "HawAPI", "ghcr.io/hawapi/hawapi", "v1.0.0"},
	}
	ID = 5
)

func main() {
	tmpl := template.Must(template.ParseFiles("index.tmpl"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, containers)
		if err != nil {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("GET /api/containers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(containers)
	})

	http.HandleFunc("POST /api/containers", func(w http.ResponseWriter, r *http.Request) {
		var c Container
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		ID++
		c.ID = ID
		containers = append(containers, c)

		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("DELETE /api/containers/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		for i, c := range containers {
			if c.ID == id {
				containers = append(containers[:i], containers[i+1:]...)
				break
			}
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
