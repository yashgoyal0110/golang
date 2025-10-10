package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) Load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &tm.Tasks)
}

func (tm *TaskManager) Save(filename string) error {
	data, err := json.MarshalIndent(tm.Tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (tm *TaskManager) AddTask(title, description string) Task {
	id := 1
	if len(tm.Tasks) > 0 {
		id = tm.Tasks[len(tm.Tasks)-1].ID + 1
	}
	task := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	tm.Tasks = append(tm.Tasks, task)
	return task
}

func (tm *TaskManager) CompleteTask(id int) bool {
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks[i].Completed = true
			return true
		}
	}
	return false
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}

func main() {
	filename := "tasks.json"
	manager := &TaskManager{}
	manager.Load(filename)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tasks := manager.ListTasks()
			json.NewEncoder(w).Encode(tasks)
			return
		}
		if r.Method == http.MethodPost {
			var input struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}
			json.NewDecoder(r.Body).Decode(&input)
			task := manager.AddTask(input.Title, input.Description)
			manager.Save(filename)
			json.NewEncoder(w).Encode(task)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/tasks/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var input struct {
			ID int `json:"id"`
		}
		json.NewDecoder(r.Body).Decode(&input)
		ok := manager.CompleteTask(input.ID)
		manager.Save(filename)
		if ok {
			fmt.Fprint(w, "Task completed")
		} else {
			http.Error(w, "Task not found", http.StatusNotFound)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
