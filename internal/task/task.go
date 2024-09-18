package task

// Task structure
type Task struct {
	Label string
	Id    int
}

// Create a new task with the given label
func Create(label string) (t *Task) {
	t = &Task{
		Label: label,
	}
	return
}
