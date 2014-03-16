package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func ListTasks(r render.Render) {
	r.JSON(200, []string{"one", "two", "three"})
}

func GetTask(r render.Render, params martini.Params) {
	r.JSON(200, map[string]interface{}{params["task"]: "world"})
}
