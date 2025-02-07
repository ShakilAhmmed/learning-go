package database

type Model struct {
	DatabaseContract
}

func NewModel() Model {
	return Model{DatabaseContract: Driver}
}
