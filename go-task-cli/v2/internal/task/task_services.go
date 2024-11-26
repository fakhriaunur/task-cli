package task

type TaskService struct {
	repo taskRepoPort
}

func NewTaskService(r taskRepoPort) TaskServicePort {
	return &TaskService{
		repo: r,
	}
}

// add implements TaskRepoPort.
func (t *TaskService) AddTask(desc string) (task, error) {
	return t.repo.add(desc)
}

// delete implements TaskRepoPort.
func (t *TaskService) DeleteTask(id int) error {
	return t.repo.delete(id)
}

// listAll implements TaskRepoPort.
func (t *TaskService) ListAllTasks() ([]task, error) {
	return t.repo.listAll()
}

// listByStatus implements TaskRepoPort.
func (t *TaskService) ListTasksByStatus(status string) ([]task, error) {
	switch status {
	case statusDone, statusInProgress, statusTodo:
		return t.repo.listByStatus(status)
	}

	return nil, nil
}

// markDone implements TaskRepoPort.
func (t *TaskService) MarkTaskDone(id int) error {
	return t.repo.markDone(id)
}

// markInProgress implements TaskRepoPort.
func (t *TaskService) MarkTaskInProgress(id int) error {
	return t.repo.markInProgress(id)
}

// update implements TaskRepoPort.
func (t *TaskService) UpdateTask(id int, desc string) error {
	return t.repo.update(id, desc)
}
