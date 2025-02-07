package main

import "fmt"

type PgSQL struct {
	rawQuery string
	table    string
}

func (database *PgSQL) Query(table string) DatabaseContract {
	database.table = table
	database.rawQuery = ""
	return database
}

func (database *PgSQL) Where(column, value string) DatabaseContract {
	if database.table == "" {
		panic("Table Name is Empty")
	}
	database.rawQuery += "SELECT * FROM " + database.table + " WHERE " + column + "=\"" + value + "\""
	return database
}

func (database *PgSQL) Get() {
	fmt.Println(database.rawQuery)
}
