package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    r := mux.NewRouter()

    s1 := r.PathPrefix("/path1").Subrouter()
    s1.HandleFunc("/", Handler1)

    s2 := r.PathPrefix("/path2").Subrouter()
    s2.HandleFunc("/", Handler2)

    log.Println("Listening ...")
    http.ListenAndServe(":8080", r)
}

func Handler1(resp http.ResponseWriter, _ *http.Request) {

    fmt.Fprint(resp, "Subroute 1")
}

func Handler2(resp http.ResponseWriter, _ *http.Request) {

    fmt.Fprint(resp, "Subroute 2")
}