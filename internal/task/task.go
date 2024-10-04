package task

import (
	"time"
)

type Task struct {
	Id        int
	Label     string
	CreatedAt time.Time
	DoneAt    time.Time
}

func CreateTask(label string) (t *Task) {
	t = &Task{
		Id:        1,
		Label:     label,
		CreatedAt: time.Now(),
	}
	return
}

func (t *Task) MarkAsDone() time.Time {
	t.DoneAt = time.Now()
	return t.DoneAt
}
