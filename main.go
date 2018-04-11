package main

import (
  "os/exec"
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
  out, err := exec.Command("/s3x","ls","s3://capnchainsaw-test").Output()
  if err != nil {
    log.WithFields(log.Fields{
      "error":    err,
    }).Error("Failed query files!")
  } else {
    listTxt := string(out)
    // TODO split string
    indexGuts.Files = append(indexGuts.Files, listTxt)
  }

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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
  // Download file local from Upload
  // TODO

  // Copy to bucket
  // TODO

  // Delete local files
  // TODO

  DisplayIndex(w, r)
}

func main() {
  s3Mux := mux.NewRouter()
  s3Mux.HandleFunc("/", DisplayIndex).Methods("GET")
  s3Mux.HandleFunc("/upload", UploadHandler).Methods("POST")
  http.Handle("/", s3Mux)

  http.ListenAndServe(":8080", nil)
}
