package project

import (
	"time"

	"github.com/florian-renfer/tasks/internal/task"
)

type Project struct {
	Id          int
	Label       string
	SubProjects []*Project
	Tasks       []*task.Task
	CreatedAt   time.Time
}

func CreateProject(label string) (p *Project) {
	p = &Project{
		Id:          1,
		Label:       label,
		SubProjects: make([]*Project, 0),
		Tasks:       make([]*task.Task, 0),
		CreatedAt:   time.Now(),
	}
	return
}

func (p *Project) AddSubProject(label string) (c *Project) {
	c = &Project{
		Id:          1,
		Label:       label,
		SubProjects: make([]*Project, 0),
		Tasks:       make([]*task.Task, 0),
		CreatedAt:   time.Now(),
	}
	p.SubProjects = append(p.SubProjects, c)
	return
}

func (p *Project) AddTask(label string) (t *task.Task) {
	t = task.CreateTask(label)
	p.Tasks = append(p.Tasks, t)
	return
}
