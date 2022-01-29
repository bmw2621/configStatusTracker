package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func saveConfig(w http.ResponseWriter, r *http.Request) {
	yaml, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	} 

	f, err := os.Create(fmt.Sprintf("./configs/%v.yaml", time.Now().Unix()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	defer f.Close()

	_, err = f.Write(yaml)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success"))
	
}

func getConfigs (w http.ResponseWriter, r *http.Request){
	files, err := os.ReadDir("./configs")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}

	var filesIds []string;
	for _, f := range files {
		filename := f.Name()
		filesIds = append(filesIds, filename[:len(filename)-5])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filesIds)
}

func getConfig(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	dat, err := os.ReadFile(fmt.Sprintf("./configs/%v.yaml",params["id"]))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(dat)
}
