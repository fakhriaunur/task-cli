package task

type TaskServicePort interface {
	AddTask(desc string) (task, error)
	UpdateTask(id int, status string) error
	DeleteTask(id int) error
	MarkTaskInProgress(id int) error
	MarkTaskDone(id int) error
	ListAllTasks() ([]task, error)
	ListTasksByStatus(status string) ([]task, error)
}

type TaskRepoConfigPort[T any] interface {
	write(T) error
	load() (T, error)
	reset() error
}

type taskRepoPort interface {
	add(desc string) (task, error)
	update(id int, status string) error
	delete(id int) error
	markInProgress(id int) error
	markDone(id int) error
	listAll() ([]task, error)
	listByStatus(status string) ([]task, error)
}

type taskMapperPort interface {
	encode(taskStructure) (dat []byte, err error)
	decode(dat []byte) (taskStructure, error)
}

type taskDAOPort interface {
	saveTaskStructure(taskStructure) error
	loadTaskStructure() (taskStructure, error)
}
