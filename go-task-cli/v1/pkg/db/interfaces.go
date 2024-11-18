package db

type InMemoryPort[T any] interface {
	Write(T) error
	Load() (T, error)
	CommonDB
}

type InMemoryRepo interface {
	Write() error
	Load() error
	CommonDB
}

type EngineBasedDB[T any] interface {
	Connect() error
	Close() error
	InMemoryPort[T]
}

type Database[T any] interface {
	EngineBasedDB[T]
	InMemoryPort[T]
}

type CommonDB interface {
	Reset() error
}
