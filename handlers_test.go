package main

import (
	"html/template"
	"testing"

	"log"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

type MockRenderer struct {
	status int
	data   interface{}
}

func (m *MockRenderer) JSON(status int, v interface{}) {
	m.status = status
	m.data = v
}

func (m *MockRenderer) HTML(status int, name string, v interface{}, htmlOpt ...render.HTMLOptions) {
	return
}
func (m *MockRenderer) Error(status int) {
	return
}
func (m *MockRenderer) Redirect(location string, status ...int) {
	return
}
func (m *MockRenderer) Template() *template.Template {
	return nil
}

var fExpect = "Expected %#v but got %#v"

func TestListTasks(t *testing.T) {
	db := initDb("test.db")
	defer db.DropTables()

	// Nothing in the database
	mRender := &MockRenderer{}
	log := &log.Logger{}
	ListTasks(mRender, db, log)
	if mRender.status != 200 {
		t.Errorf("Expected %#v but got %#v", 200, mRender.status)
		return
	}
	actual := mRender.data.([]tasksView)
	if len(actual) != 0 {
		t.Errorf("Expected %#v but got %#v", 0, len(actual))
		return
	}

	// One task in the database
	task := &Task{Name: "MyName", Script: "MyScript"}
	db.Insert(task)

	ListTasks(mRender, db, log)
	if mRender.status != 200 {
		t.Errorf("Expected %#v but got %#v", 200, mRender.status)
		return
	}
	expected := tasksView{Id: 1, Name: "MyName"}
	actual = mRender.data.([]tasksView)
	if len(actual) != 1 || actual[0] != expected {
		t.Errorf("Expected %#v but got %#v", expected, actual)
		return
	}
}

func TestGetTask(t *testing.T) {
	db := initDb("test.db")
	defer db.DropTables()

	task := &Task{Name: "MyName", Script: "MyScript"}
	db.Insert(task)

	// bad parameter
	mRender := &MockRenderer{}
	params := martini.Params{"id": "not a number"}
	GetTask(mRender, params, db)
	if mRender.status != 400 {
		t.Errorf(fExpect, 400, mRender.status)
		return
	}

	// ok
	params = martini.Params{"id": "1"}
	GetTask(mRender, params, db)
	if mRender.status != 200 {
		t.Errorf(fExpect, 200, mRender.status)
		return
	}
	actual := mRender.data.(*Task)
	if *actual != *task {
		t.Errorf(fExpect, task, actual)
		return
	}

	// not found
	params = martini.Params{"id": "99"}
	GetTask(mRender, params, db)
	if mRender.status != 404 {
		t.Errorf(fExpect, 404, mRender.status)
		return
	}
}

func TestAddTask(t *testing.T) {
	db := initDb("test.db")
	defer db.DropTables()

	mRender := &MockRenderer{}
	// Note you cannot set your own Id
	payload := Task{Id: 5, Name: "Hello", Script: "echo 'hello'"}
	AddTask(mRender, payload, db)

	if mRender.status != 201 {
		t.Errorf(fExpect, 201, mRender.status)
		return
	}
	actual := mRender.data.(Task)
	payload.Id = 1
	if actual != payload {
		t.Errorf(fExpect, payload, actual)
		return
	}

	count, _ := db.SelectInt("select count(*) from tasks")
	if count != 1 {
		t.Errorf(fExpect, 1, count)
		return
	}
}

func TestUpdateTask(t *testing.T) {
	db := initDb("test.db")
	defer db.DropTables()

	task := &Task{Name: "Test", Script: "drop database"}
	db.Insert(task)

	// ok
	mRender := &MockRenderer{}
	params := martini.Params{"id": "1"}
	// Note you cannot change the Id
	payload := Task{Id: 9, Name: "Testing", Script: "updated"}
	UpdateTask(mRender, params, payload, db)

	if mRender.status != 200 {
		t.Errorf(fExpect, 200, mRender.status)
		return
	}
	actual := mRender.data.(Task)
	payload.Id = 1
	if actual != payload {
		t.Errorf(fExpect, payload, actual)
		return
	}

	// updating something that does not exist
	params = martini.Params{"id": "99"}
	UpdateTask(mRender, params, payload, db)

	if mRender.status != 404 {
		t.Errorf(fExpect, 404, mRender.status)
		return
	}

	// param is not an int
	params = martini.Params{"id": "hi"}
	UpdateTask(mRender, params, payload, db)

	if mRender.status != 400 {
		t.Errorf(fExpect, 400, mRender.status)
		return
	}
}

func TestDeleteTask(t *testing.T) {
	db := initDb("test.db")
	defer db.DropTables()

	task := &Task{Name: "Test", Script: "drop database"}
	db.Insert(task)

	// ok
	mRender := &MockRenderer{}
	params := martini.Params{"id": "1"}
	DeleteTask(mRender, params, db)
	if mRender.status != 200 {
		t.Errorf(fExpect, 200, mRender.status)
		return
	}

	count, _ := db.SelectInt("select count(*) from tasks")
	if count != 0 {
		t.Errorf(fExpect, 0, count)
		return
	}

	// id is not an integer
	params = martini.Params{"id": "hi"}
	DeleteTask(mRender, params, db)
	if mRender.status != 400 {
		t.Errorf(fExpect, 400, mRender.status)
		return
	}
}
