package tasktracker

import "github.com/fakhriaunur/task-cli/pkg/db"

type AppRepoPort interface {
	TaskRepoPort
	UserRepoPort
}

type AppRepoInMemoryPort[T any] interface {
	db.InMemoryPort[T]
	AppRepoPort
}

type UserRepoPort interface {
	AddUser(username string) error
	DelelteUserByID(id int) error
	UpdateUser(id int) error
}

type TaskRepoPort interface {
	Add(desc string) (Task, error)
	Update(id int, status string) error
	Delete(id int) error
	Mark(id int, status string) error
	ListAll() ([]Task, error)
	List(status string) ([]Task, error)
}

type TaskRepoCommonPort interface {
	TaskRepoPort
	db.CommonDB
}

type TaskRepoInMemoryPort[T any] interface {
	db.InMemoryPort[T]
	TaskRepoPort
}

type TaskRepoEngineBasedPort[T any] interface {
	db.EngineBasedDB[T]
	TaskRepoPort
}

type TaskRepoInMemory interface {
	db.InMemoryRepo
	TaskRepoPort
}
