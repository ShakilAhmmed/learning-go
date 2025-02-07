package database

var Driver DatabaseContract

const (
	MYSQL = "mysql"
	PGSQL = "pgsql"
)

func InitDatabase(dbType string) {
	switch dbType {
	case MYSQL:
		Driver = MySQL()
	case PGSQL:
		Driver = MySQL()
	default:
		panic("Unsupported Driver")
	}
}
