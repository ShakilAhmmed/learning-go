package database

import (
	"fmt"
)

type mySQL struct {
	rawQuery string
	table    string
}

func (database *mySQL) SetTable(table string) {
	database.table = table
}

func (database *mySQL) Query() DatabaseContract {
	database.rawQuery = ""
	return database
}

func (database *mySQL) Where(column, value string) DatabaseContract {
	if database.table == "" {
		panic("Table Name is Empty")
	}
	database.rawQuery += "SELECT * FROM " + database.table + " WHERE " + column + "=\"" + value + "\""
	return database
}

func (database *mySQL) Get() {
	fmt.Println(database.rawQuery)
}

func MySQL() *mySQL {
	return new(mySQL)
}
