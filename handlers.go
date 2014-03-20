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
	err := db.Insert(&taskPayload)
	if err != nil {
		log.Printf("Error inserting: %v", err)
		r.JSON(400, map[string]string{"message": "failed inserting task"})
		return
	}
	r.JSON(201, taskPayload)
}

func UpdateTask(r render.Render, params martini.Params, taskPayload Task, db *gorp.DbMap) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.JSON(400, map[string]string{"message": "id must be an integer"})
		return
	}
	taskPayload.Id = id
	count, err := db.Update(&taskPayload)
	if count == 0 {
		r.JSON(404, map[string]string{"message": "task not found"})
		return
	}
	if err != nil {
		log.Printf("Failed updating task %v: %v", id, err)
		r.JSON(500, map[string]string{"message": "Failed to update task"})
		return
	}
	r.JSON(200, taskPayload)
}

func DeleteTask(r render.Render, params martini.Params, db *gorp.DbMap) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.JSON(400, map[string]string{"message": "id must be an integer"})
		return
	}
	count, err := db.Delete(&Task{Id: id})
	if err != nil {
		log.Printf("Failed deleting task %v: %v", id, err)
		r.JSON(500, map[string]string{"messaage": "Failed to delete task"})
		return
	}
	if count != 1 {
		r.JSON(404, map[string]string{"message": "Task not found"})
		return
	}
	r.JSON(200, map[string]string{})
}
