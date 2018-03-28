package main

import (
  "html/template"
  "net/http"

  "github.com/gorilla/mux"
  log "github.com/Sirupsen/logrus"
)

type IndexPage struct {
  Files []string
}

func DisplayIndex(w http.ResponseWriter, r *http.Request) {
  indexGuts := IndexPage{
    Files: []string{},
  }

  // Query S3 for file list.
  // TODO

  // Populate template and display.
  t, err := template.ParseFiles("index.html")
  if err != nil {
    log.WithFields(log.Fields{
      "error":    err,
    }).Error("Failed to parse templates!")
    return
  }
  t.Execute(w, indexGuts)
}

func main() {
  s3Mux := mux.NewRouter()
  s3Mux.HandleFunc("/", DisplayIndex).Methods("GET")
  http.Handle("/", s3Mux)

  http.ListenAndServe(":8080", nil)
}
