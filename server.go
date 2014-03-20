package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/tasks", ListTasks)
	m.Get("/tasks/:id", GetTask)

	m.Map(initDb())

	m.Run()
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		panic("sql.Open failed")
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'tasks' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Task{}, "tasks").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		panic("Create tables failed")
	}
	return dbmap
}
