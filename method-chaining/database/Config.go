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
		Driver = PgSQL()
	default:
		panic("Unsupported Driver")
	}
}
