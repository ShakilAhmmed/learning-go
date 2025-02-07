package database

type DatabaseContract interface {
	SetTable(table string)
	Query() DatabaseContract
	Where(column, value string) DatabaseContract
	Get()
}
