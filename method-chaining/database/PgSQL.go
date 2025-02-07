package database

import (
	"fmt"
)

type pgSQL struct {
	rawQuery string
	table    string
}

func (database *pgSQL) SetTable(table string) {
	database.table = table
}

func (database *pgSQL) Query() DatabaseContract {
	database.rawQuery = ""
	return database
}

func (database *pgSQL) Where(column, value string) DatabaseContract {
	if database.table == "" {
		panic("Table Name is Empty")
	}
	database.rawQuery += "SELECT * FROM " + database.table + " WHERE " + column + "=\"" + value + "\""
	return database
}

func (database *pgSQL) Get() {
	fmt.Println(database.rawQuery)
}

func PgSQL() *pgSQL {
	return new(pgSQL)
}
