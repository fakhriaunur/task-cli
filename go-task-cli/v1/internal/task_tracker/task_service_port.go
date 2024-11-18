package tasktracker

type TaskServicePort interface {
	Add(desc string) (Task, error)
	Update(id int, desc string) error
	Delete(id int) error
	MarkInProgress(id int) error
	MarkDone(id int) error
	ListAll() ([]Task, error)
	ListByStatus(status string) ([]Task, error)
}

type TaskServiceCommonPort interface {
	TaskServicePort
	CommonServicePort
}

type CommonServicePort interface {
	HelpServicePort
	ResetServicePort
}

type HelpServicePort interface {
	Help() error
}

type ResetServicePort interface {
	Reset() error
}
