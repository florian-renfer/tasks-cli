package project_test

import (
	"testing"

	"github.com/florian-renfer/tasks/internal/project"
	"github.com/florian-renfer/tasks/internal/task"
)

func TestCreateProject(t *testing.T) {
	p := project.Create("test")

	if p.Name != "test" {
		t.Fatalf("Project name %v does not match the expected value of 'test'.", p.Name)
	}
}

func TestAddTask(t *testing.T) {
	p := project.Create("New project")
	task := task.Create("New task")
	task.Id = 1

	_, err := p.AddTask(task)

	if err != nil {
		// handle/catch error
	}
}

// TODO: RF implement
func TestAddTaskNewFile(t *testing.T) {}

// TODO: RF implement
func TestAddTaskFileLocked(t *testing.T) {}

// TODO: RF implement
func TestAddTaskExistingFile(t *testing.T) {}
