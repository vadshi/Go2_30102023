package taskstore

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id	int `json:"id"`
	Text string `json:"text"`
	Tags []string `json:"tags"`
	Due time.Time `json:"due"`
}

// TaskStore is a simple in-memory database of tasks;
type TaskStore struct {
	sync.Mutex

	tasks map[int]Task
	nextId int
}

// TaskStore constructor
func New() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

// CreateTask create a new task in the store
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id: ts.nextId,
		Text: text,
		Due: due}
	task.Tags = make([]string, len(tags))
	copy(task.Tags, tags)
	// Сохранили task в TaskStore
	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}


// GetTask retrieves a task from taskstore by id
func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	t, ok := ts.tasks[id]
	if ok {
		return t, nil
	} else {
		return Task{}, fmt.Errorf("task with id=%d not found", id)
	}

}

// GetAllTask retrieves all task from taskstore 
func (ts *TaskStore) GetAllTask() []Task {
	ts.Lock()
	defer ts.Unlock()

	allTasks := make([]Task, 0, len(ts.tasks))
	for _, task := range ts.tasks{
		allTasks = append(allTasks, task)
	}

	return allTasks

}