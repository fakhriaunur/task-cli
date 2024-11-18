package task

type TaskService struct {
	repo TaskRepoPort
}

func NewTaskService(r TaskRepoPort) TaskRepoPort {
	return &TaskService{
		repo: r,
	}
}

// add implements TaskRepoPort.
func (t *TaskService) add(desc string) (task, error) {
	return t.repo.add(desc)
}

// delete implements TaskRepoPort.
func (t *TaskService) delete(id int) error {
	return t.repo.delete(id)
}

// listAll implements TaskRepoPort.
func (t *TaskService) listAll() ([]task, error) {
	return t.repo.listAll()
}

// listByStatus implements TaskRepoPort.
func (t *TaskService) listByStatus(status string) ([]task, error) {
	switch status {
	case statusDone, statusInProgress, statusTodo:
		return t.repo.listByStatus(status)
	}

	return nil, nil
}

// markDone implements TaskRepoPort.
func (t *TaskService) markDone(id int, status string) error {
	return t.markDone(id, statusDone)
}

// markInProgress implements TaskRepoPort.
func (t *TaskService) markInProgress(id int, status string) error {
	return t.markInProgress(id, statusInProgress)
}

// update implements TaskRepoPort.
func (t *TaskService) update(id int) error {
	return t.update(id)
}
