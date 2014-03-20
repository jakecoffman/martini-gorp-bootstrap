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

	m.Map(initDb("dev.db"))

	m.Run()
}

func initDb(name string) *gorp.DbMap {
	db, err := sql.Open("sqlite3", name)
	nilOrPanic(err)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	err = dbmap.DropTables()
	nilOrPanic(err)
	dbmap.AddTableWithName(Task{}, "tasks").SetKeys(true, "Id")

	// TODO: Use DB migration tool
	err = dbmap.CreateTablesIfNotExists()
	nilOrPanic(err)

	return dbmap
}

func nilOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
