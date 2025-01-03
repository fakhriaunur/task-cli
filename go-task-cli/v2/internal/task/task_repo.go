package task

import (
	"log"
	"os"
)

type TaskRepo struct {
	dao taskDAOPort
}

func NewTaskRepo(dao taskDAOPort) (taskRepoPort, error) {
	repo := &TaskRepo{
		dao: dao,
	}

	return repo, nil
}

// add implements TaskRepoPort.
func (r *TaskRepo) add(desc string) (task, error) {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		log.Println("here load")
		return task{}, err
	}

	newID := len(ts.Tasks)
	newTask := task{
		ID:          newID,
		Status:      statusTodo,
		Description: desc,
		CreatedAt:   timeNow,
		UpdatedAt:   timeZero,
	}

	ts.Tasks[newID] = newTask

	if err := r.dao.saveTaskStructure(ts); err != nil {
		return task{}, err
	}

	return newTask, nil
}

// delete implements TaskRepoPort.
func (r *TaskRepo) delete(id int) error {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		return err
	}

	_, ok := ts.Tasks[id]
	if !ok {
		return os.ErrNotExist
	}

	delete(ts.Tasks, id)

	if err := r.dao.saveTaskStructure(ts); err != nil {
		return err
	}

	return nil
}

// listAll implements TaskRepoPort.
func (r *TaskRepo) listAll() ([]task, error) {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		return nil, err
	}

	var tasks []task
	for _, task := range ts.Tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// listByStatus implements TaskRepoPort.
func (r *TaskRepo) listByStatus(status string) ([]task, error) {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		return nil, err
	}

	var tasks []task
	for _, task := range ts.Tasks {
		if task.Status == status {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// markDone implements TaskRepoPort.
func (r *TaskRepo) markDone(id int) error {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		return err
	}

	task, ok := ts.Tasks[id]
	if !ok {
		return os.ErrNotExist
	}

	task.Status = statusDone
	ts.Tasks[id] = task

	if err := r.dao.saveTaskStructure(ts); err != nil {
		return err
	}

	return nil
}

// markInProgress implements TaskRepoPort.
func (r *TaskRepo) markInProgress(id int) error {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		return err
	}

	task, ok := ts.Tasks[id]
	if !ok {
		return os.ErrNotExist
	}

	task.Status = statusInProgress
	ts.Tasks[id] = task

	if err := r.dao.saveTaskStructure(ts); err != nil {
		return err
	}

	return nil
}

// update implements TaskRepoPort.
func (r *TaskRepo) update(id int, desc string) error {
	ts, err := r.dao.loadTaskStructure()
	if err != nil {
		return err
	}

	task, ok := ts.Tasks[id]
	if !ok {
		return os.ErrNotExist
	}

	task.Description = desc
	ts.Tasks[id] = task

	if err := r.dao.saveTaskStructure(ts); err != nil {
		return err
	}

	return nil
}

var _ taskRepoPort = (*TaskRepo)(nil)
