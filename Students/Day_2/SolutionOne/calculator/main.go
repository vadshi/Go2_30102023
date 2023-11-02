package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var numberOne int
var numberTwo int

type First struct {
	First int `json:"first"`
}

type Second struct {
	First int `json:"first"`
}

type Res struct {
	First  int    `json:"first"`
	Second int    `json:"second"`
	Result string `json:"result"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hello form calculator web-server 
	</h1><br><br><h1>API:</h1><br><h2>/info - this information
	</h2><br><h2>/first - generate first random number</h2>
	<br><h2>/second - generate second random number</h2><br>
	<h2>/add - first + second</h2><br><h2>/sub - first - second</h2>
	<br><h2>/mul - first * second</h2><br><h2>/div - first / second</h2><br>`)
}

func getFirst(w http.ResponseWriter, r *http.Request) {
	numberOne = rand.Intn(100)
	w.Header().Set("Content-Type", "application/json")
	f := First{numberOne}
	json.NewEncoder(w).Encode(f)
}

func getSecond(w http.ResponseWriter, r *http.Request) {
	numberTwo = rand.Intn(100)
	w.Header().Set("Content-Type", "application/json")
	f := Second{numberTwo}
	json.NewEncoder(w).Encode(f)
}

func add(w http.ResponseWriter, r *http.Request) {
	res := numberOne + numberTwo
	w.Header().Set("Content-Type", "application/json")
	var f Res
	f.First = numberOne
	f.Second = numberTwo
	f.Result = fmt.Sprintf("%d+%d=%d", numberOne, numberTwo, res)
	json.NewEncoder(w).Encode(f)
}

func sub(w http.ResponseWriter, r *http.Request) {
	res := numberOne - numberTwo
	w.Header().Set("Content-Type", "application/json")
	var f Res
	f.First = numberOne
	f.Second = numberTwo
	f.Result = fmt.Sprintf("%d-%d=%d", numberOne, numberTwo, res)
	json.NewEncoder(w).Encode(f)
}

func mul(w http.ResponseWriter, r *http.Request) {
	res := numberOne * numberTwo
	w.Header().Set("Content-Type", "application/json")
	var f Res
	f.First = numberOne
	f.Second = numberTwo
	f.Result = fmt.Sprintf("%d*%d=%d", numberOne, numberTwo, res)
	json.NewEncoder(w).Encode(f)
}

func div(w http.ResponseWriter, r *http.Request) {
	res := float64(numberOne) / float64(numberTwo)
	w.Header().Set("Content-Type", "application/json")
	var f Res
	f.First = numberOne
	f.Second = numberTwo
	f.Result = fmt.Sprintf("%d/%d=%f", numberOne, numberTwo, res)
	json.NewEncoder(w).Encode(f)
}

func main() {
	fmt.Println("Starting web server")
	// Changed from version 1.20
	rand.New(rand.NewSource(time.Now().UnixNano()))


	http.HandleFunc("/", GetInfo)
	http.HandleFunc("/info/", GetInfo)
	http.HandleFunc("/first/", getFirst)
	http.HandleFunc("/second/", getSecond)
	http.HandleFunc("/add/", add)
	http.HandleFunc("/sub/", sub)
	http.HandleFunc("/mul/", mul)
	http.HandleFunc("/div/", div)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
