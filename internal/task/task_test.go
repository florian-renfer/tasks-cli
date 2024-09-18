package task_test

import (
	"testing"

	"github.com/florian-renfer/tasks/internal/task"
)

func TestCreateTask(t *testing.T) {
	l := "Task 1"
	c := task.Create(l)
	if l != c.Label {
		t.Fatalf("Expected %v but got %v.", l, c.Label)
	}
}

func TestCreateMultipleTasks(t *testing.T) {
	a := task.Create("1")
	b := task.Create("2")
	c := task.Create("3")
	tasks := []*task.Task{a, b, c}
	if len(tasks) != 3 {
		t.Fatalf("Expected a length of %d but got %v.", 3, len(tasks))
	}
}
