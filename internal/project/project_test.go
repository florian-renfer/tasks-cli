package project_test

import (
	"testing"

	"github.com/florian-renfer/tasks/internal/project"
)

func TestCreateProject(t *testing.T) {
	p := project.CreateProject("Project label")
	want := "Project label"
	if p.Label != want {
		t.Fatalf("Label for created project does not meet the expectation. (Got: %v, Want: %v).\n", p.Label, want)
	}
}

func TestAddSubProject(t *testing.T) {
	p := project.CreateProject("Project label")
	if len(p.SubProjects) != 0 {
		t.Fatalf("Project should not contain any sub-projects by now, but got %d sub-projects instead.\n", len(p.SubProjects))
	}

	want := "Sub-project label"
	s := p.AddSubProject("Sub-project label")
	if len(p.SubProjects) != 1 {
		t.Fatalf("Project should contain 1 sub-project by now, but got %d tasks instead.\n", len(p.SubProjects))
	}
	if s.Label != want {
		t.Fatalf("Expected sub-project label was %v, but got %v instead.", want, s.Label)
	}
}

func TestAddTask(t *testing.T) {
	p := project.CreateProject("Project label")
	if len(p.Tasks) != 0 {
		t.Fatalf("Project should not contain any tasks by now, but got %d tasks instead.\n", len(p.Tasks))
	}

	want := "Task label"
	c := p.AddTask("Task label")
	if len(p.Tasks) != 1 {
		t.Fatalf("Project should contain 1 task by now, but got %d tasks instead.\n", len(p.Tasks))
	}
	if c.Label != want {
		t.Fatalf("Expected task label was %v, but got %v instead.", want, c.Label)
	}
}

func TestAddTaskToSubProject(t *testing.T) {
	p := project.CreateProject("Project label")
	if len(p.SubProjects) != 0 {
		t.Fatalf("Project should not contain any sub-projects by now, but got %d sub-projects instead.\n", len(p.SubProjects))
	}

	want := "Sub-project label"
	s := p.AddSubProject("Sub-project label")
	if len(p.SubProjects) != 1 {
		t.Fatalf("Project should contain 1 sub-project by now, but got %d tasks instead.\n", len(p.SubProjects))
	}
	if s.Label != want {
		t.Fatalf("Expected sub-project label was %v, but got %v instead.", want, s.Label)
	}

	want = "Sub-project task label"
	c := s.AddTask("Sub-project task label")
	if len(p.SubProjects) != 1 {
		t.Fatalf("Project should not contain any sub-projects by now, but got %d sub-projects instead.\n", len(p.SubProjects))
	}
	if len(s.Tasks) != 1 {
		t.Fatalf("Sub-project should contain 1 task by now, but got %d tasks instead.\n", len(s.Tasks))
	}
	if c.Label != want {
		t.Fatalf("Expected task label was %v, but got %v instead.", want, c.Label)
	}
}
