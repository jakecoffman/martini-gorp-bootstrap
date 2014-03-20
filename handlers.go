package main

import (
	"log"
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/martini-contrib/render"
)

type ListTasksView struct {
	Id   int
	Name string
}

func ListTasks(r render.Render, db *gorp.DbMap, log *log.Logger) {
	var taskIds []ListTasksView
	_, err := db.Select(&taskIds, "select id,name from tasks order by id")
	if err != nil {
		log.Printf("Error selecting from database: %v", err)
		r.JSON(500, map[string]string{"message": "error while retrieving tasks"})
		return
	}
	r.JSON(200, taskIds)
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
