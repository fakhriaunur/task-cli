package task

type TaskServicePort interface {
	AddTask(desc string) (task, error)
	UpdateTask(id int) error
	DeleteTask(id int) error
	MarkTask(id int) error
	ListTasks(func(*string) ([]task, error)) ([]task, error)
}

type TaskRepoConfig[T any] interface {
	write(T) error
	load() (T, error)
	reset()
}

type TaskRepoPort interface {
	add(desc string) (task, error)
	update(id int) error
	delete(id int) error
	markInProgress(id int, status string) error
	markDone(id int, status string) error
	listAll() ([]task, error)
	listByStatus(status string) ([]task, error)
}
