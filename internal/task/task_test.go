package task_test

import (
	"os"
	"testing"

	"github.com/florian-renfer/tasks/internal/task"
)

func TestCreateTask(t *testing.T) {
	got := task.CreateTask("Label")
	want := "Label"
	if got.Label != want {
		t.Fatalf("Label for created task does not meet the expectation. (Got: %v, Want: %v).\n", got.Label, want)
	}
}

func TestPrint(t *testing.T) {
	c := task.CreateTask("Label")
	c.Print(os.Stdout)
}
