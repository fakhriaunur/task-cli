package tasktracker

import "fmt"

type TaskService struct {
	Repo TaskRepoCommonPort
}

var _ TaskServiceCommonPort = (*TaskService)(nil)

func NewTaskService(repo TaskRepoCommonPort) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (r *TaskService) Add(desc string) (Task, error) {
	return r.Repo.Add(desc)
}

func (r *TaskService) Delete(id int) error {
	return r.Repo.Delete(id)
}

func (r *TaskService) Update(id int, desc string) error {
	return r.Repo.Update(id, desc)
}

func (r *TaskService) MarkInProgress(id int) error {
	return r.Repo.Mark(id, StatusInProgress)
}

func (r *TaskService) MarkDone(id int) error {
	return r.Repo.Mark(id, StatusDone)
}

func (r *TaskService) ListAll() ([]Task, error) {
	return r.Repo.ListAll()
}

func (r *TaskService) ListByStatus(status string) ([]Task, error) {
	return r.Repo.List(status)
}

func (r *TaskService) Help() error {
	fmt.Println("Usage etc...")
	return nil
}

func (r *TaskService) Reset() error {
	return r.Repo.Reset()
}
