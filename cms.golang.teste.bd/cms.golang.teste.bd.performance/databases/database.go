package database

type IDatabase interface {
	StartDB()
	GetDatabase() interface{}
	Close()
}
