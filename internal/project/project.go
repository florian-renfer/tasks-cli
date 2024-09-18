package project

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/florian-renfer/tasks/internal/task"
)

// Project structure
type Project struct {
	Name  string
	Tasks []task.Task
}

// Create a new project with the given name
func Create(name string) (p *Project) {
	p = &Project{
		Name: name,
	}
	return
}

// Add the given task to the current project
func (p *Project) AddTask(task *task.Task) (t *task.Task, err error) {

	// Validate task and project, e.g. task has to be unique for the given project, e.g. ignore case when comparing task labels
	f, err := os.OpenFile("../../tasks-cli.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to create file.\n%v\n", err)
	}

	// Close .csv file right before exiting the function
	defer f.Close()

	w := csv.NewWriter(f)
	r := []string{strconv.Itoa(task.Id), task.Label}

	err = w.Write(r)
	if err != nil {
		log.Fatalf("Error writing to file.\n%v\n", err)
	}

	w.Flush()

	fmt.Printf("Task '%v' added to project '%v'.", task.Label, p.Name)

	t = task
	return t, nil
}
