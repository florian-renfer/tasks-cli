package task

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Task structure to store task related data
type Task struct {
	Label string
	Id    int
}

type TaskService interface {
	Create(task *Task) error
}

type CsvTaskService struct {
	tasks []*Task
}

func GetCsvTaskServiceInstance() (c *CsvTaskService) {
	c = &CsvTaskService{}
	c.tasks = make([]*Task, 0, 4)
	return
}

// Temporary store
var tasks = make([]*Task, 0, 4)

// Creates a pointer to a new task instance with the given label
func (c *CsvTaskService) Create(label string) (t *Task) {
	t = &Task{}
	t.Id = len(tasks)
	t.Label = label
	tasks = append(tasks, t)

	f, err := os.OpenFile("tasks-cli.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to create file.\n%v\n", err)
	}

	defer f.Close()

	w := csv.NewWriter(f)
	r := []string{strconv.Itoa(t.Id), t.Label}

	err = w.Write(r)
	if err != nil {
		log.Fatalf("Error writing to file.\n%v\n", err)
	}

	w.Flush()

	return
}

// Deletes a task and destroys its pointer by a given label
func Delete(label string) (id uint) {
	return 0
}

// List tasks based on their state
func List(all bool) {
	for _, t := range tasks {
		fmt.Printf("%v\t%v\n", t.Id, t.Label)
	}
}
