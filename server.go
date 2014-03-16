package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/tasks", ListTasks)
	m.Get("/tasks/:task", GetTask)

	m.Run()
}
