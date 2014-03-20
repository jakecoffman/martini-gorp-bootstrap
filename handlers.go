package main

import (
	"log"
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/martini-contrib/render"
)

func ListTasks(r render.Render, db *gorp.DbMap) {
	var tasks []Task
	_, err := db.Select(&tasks, "select * from tasks order by id")
	if err != nil {
		log.Fatalln("ERRORS")
	}
	r.JSON(200, tasks)
}

func GetTask(r render.Render, params martini.Params, db *gorp.DbMap) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.JSON(400, map[string]string{"message": "id must be an integer"})
		return
	}
	task, err := db.Get(Task{}, id)
	if err != nil {
		r.JSON(500, map[string]string{"message": "error while retrieving task"})
		return
	}
	if task == nil {
		r.JSON(404, map[string]string{"message": "task not found"})
		return
	}
	r.JSON(200, task)
}
