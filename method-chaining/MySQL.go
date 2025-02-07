package main

import (
	"fmt"
)

type MySQL struct {
	rawQuery string
	table    string
}

func (database *MySQL) Query(table string) DatabaseContract {
	database.table = table
	database.rawQuery = ""
	return database
}

func (database *MySQL) Where(column, value string) DatabaseContract {
	if database.table == "" {
		panic("Table Name is Empty")
	}
	database.rawQuery += "SELECT * FROM " + database.table + " WHERE " + column + "=\"" + value + "\""
	return database
}

func (database *MySQL) Get() {
	fmt.Println(database.rawQuery)
}
