package main

import (
	"log"
	"strconv"

	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/martini-contrib/render"
)

type tasksView struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ListTasks(r render.Render, db *gorp.DbMap, log *log.Logger) {
	var taskIds []tasksView
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

func AddTask(r render.Render, taskPayload Task, db *gorp.DbMap) {
	// TODO: Check if this accepts Ids externally. If so, remove them.
	err := db.Insert(&taskPayload)
	if err != nil {
		log.Printf("Error inserting: %v", err)
		r.JSON(400, map[string]string{"message": "failed inserting task"})
		return
	}
	r.JSON(201, taskPayload)
}

func UpdateTask(r render.Render, params martini.Params, taskPayload Task, db *gorp.DbMap) {
	// TODO: Check if this works with Ids
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.JSON(400, map[string]string{"message": "id must be an integer"})
		return
	}
	taskPayload.Id = id
	count, err := db.Update(&taskPayload)
	if err != nil || count != 1 {
		log.Printf("Failed updating task %i: %v", err)
		r.JSON(500, map[string]string{"message": "Failed to update task"})
		return
	}
	r.JSON(200, taskPayload)
}
