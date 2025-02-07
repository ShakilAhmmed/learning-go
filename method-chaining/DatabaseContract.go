package main

type DatabaseContract interface {
	Query(table string) DatabaseContract
	Where(column, value string) DatabaseContract
	Get()
}
