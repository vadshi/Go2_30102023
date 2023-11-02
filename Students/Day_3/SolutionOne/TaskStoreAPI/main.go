/*
Реализовать TaskStoreAPI с помощью GorillaMux
Реализовать в качестве хранилища SQLite
Добавить реализацию endpoints для '/tags/' и '/due/'
Важное замечание про тип time.Time и SQLite.
*/
package main

import (
	"TaskStoreAPIMod/internal/taskstore"
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling task create at %s\n", r.URL.Path)

	//структура для создания задачи
	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	//для ответа json
	type ResponseId struct {
		Id int `json:"id"`
	}

	//JSON в качестве Content-Type
	contentType := r.Header.Get("Content-Type")

	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	//обработка тела запроса
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//создали новую задачу в хранилище и получаем ее <id>
	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)

	//создаем json для ответа
	js, err := json.Marshal(ResponseId{Id: id})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	//обязательно вносим изменения в Header до вызова метода Write()
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) getAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling get all tasks at %s\n", r.URL.Path)

	allTasks := ts.store.GetAllTasks()
	js, err := json.Marshal(allTasks)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	//обязательно вносим изменения в Header до вызова метода Write()
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) getTaskHandler(w http.ResponseWriter, r *http.Request) {

	//Считаем id из строки запроса и конвертируем его в int
	vars := mux.Vars(r) // {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}

	log.Printf("Handling get task <%d> at %s\n", id, r.URL.Path)

	task, err := ts.store.GetTask(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}

	js, err := json.Marshal(task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
		return
	}

	//обязательно вносим изменения в Header до вызова метода Write()
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (ts *taskServer) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	//Считаем id из строки запроса и конвертируем его в int
	vars := mux.Vars(r) // {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}

	log.Printf("Handling delete task <%d> at %s\n", id, r.URL.Path)

	err = ts.store.DeleteTask(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}
}

func (ts *taskServer) deleteAllTaskHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Handling delete all tasks at %s\n", r.URL.Path)

	ts.store.DeleteAllTask()
}

func main() {
	server := NewTaskServer()

	log.Println("Trying to start REST API task!")

	router := mux.NewRouter()
	//1. Если на вход пришел запрос /task
	router.HandleFunc("/task/", server.createTaskHandler).Methods("POST")
	router.HandleFunc("/task/", server.getAllTaskHandler).Methods("GET")
	router.HandleFunc("/task/", server.deleteAllTaskHandler).Methods("DELETE")
	router.HandleFunc("/task/{id}", server.getTaskHandler).Methods("GET")
	router.HandleFunc("/task/{id}", server.deleteTaskHandler).Methods("DELETE")

	log.Println("Router configured successfully! Let's go!")

	log.Fatal(http.ListenAndServe(":3000", router))

}
