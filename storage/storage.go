package storage

type Storage interface {
	SaveUnhandled()
	DeleteUnhandled()
}
