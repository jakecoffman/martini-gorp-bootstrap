package main

import (
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func ListTasks(r render.Render, tasks TaskService) {
	r.JSON(200, tasks.GetAll())
}

func GetTask(r render.Render, params martini.Params, tasks TaskService) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.JSON(400, map[string]string{"message": "id must be an integer"})
		return
	}
	r.JSON(200, tasks.Get(id))
}
