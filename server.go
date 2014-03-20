package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/tasks", ListTasks)
	m.Get("/tasks/:id", GetTask)
	m.Post("/tasks", binding.Json(Task{}), AddTask)
	m.Put("/tasks/:id", binding.Json(Task{}), UpdateTask)

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
	err = dbmap.DropTables()
	nilPanic(err)
	dbmap.AddTableWithName(Task{}, "tasks").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	nilPanic(err)

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	nilPanic(err)
	return dbmap
}

func nilPanic(err error) {
	if err != nil {
		panic(err)
	}
}
