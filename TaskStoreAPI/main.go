/*
// Пример REST сервера с несколькими маршрутами(используем только стандартную библиотеку)

// POST   /task/              :  создаёт задачу и возвращает её ID
// GET    /task/<taskid>      :  возвращает одну задачу по её ID
// GET    /task/              :  возвращает все задачи
// DELETE /task/<taskid>      :  удаляет задачу по ID
// DELETE /task/              :  удаляет все задачи
// GET    /tag/<tagname>      :  возвращает список задач с заданным тегом
// GET    /due/<yy>/<mm>/<dd> :  возвращает список задач, запланированных на указанную дату

Структура проекта
https://github.com/golang-standards/project-layout/blob/master/README_ru.md
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"TaskStoreAPI/internal/taskstore"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) taskHandler(w http.ResponseWriter, r *http.Request) {
	//Request is only '/task/' URL without ID
	if r.URL.Path == "/task/" {
		if r.Method == http.MethodPost {
			ts.createTaskHandler(w, r)
		} else if r.Method == http.MethodGet{
			ts.getAllTaskHandler(w, r)
		} else if r.Method == http.MethodDelete{
			ts.deleteAllTaskHandler(w, r)
		} else {
			http.Error(w, fmt.Sprint("expect method GET, POST, DELETE at '/task', got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
		
	} else {
		// Request has an ID as '/task/<id>' URL
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2{
			http.Error(w, fmt.Sprint("expect 'task/<id>' in task handler"), http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodGet {
			ts.getTaskHandler(w, r, int(id))
		} else if r.Method == http.MethodDelete{
			ts.deleteTaskHandler(w, r, int(id))
		} else {
			http.Error(w, fmt.Sprint("expect method GET, DELETE at '/task<id>', got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}
func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()
	// Added routing for "/task/" URL
	mux.HandleFunc("/task/", server.taskHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}