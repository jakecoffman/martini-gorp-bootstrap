package main

import "testing"

func TestCRUD(t *testing.T) {
	v := tasks.GetAll()
	if len(v) != 0 {
		t.Error("Expected 0 got ", len(v))
	}

	expected1 := Task{Script: "Hi"}
	_, err := tasks.Add(&expected1)
	if err != nil {
		t.Error("Failed to add: ", err.Error())
		return
	}

	v = tasks.GetAll()
	if len(v) != 1 {
		t.Error("Expected 1 got", len(v))
		return
	}
	if *v[0] != expected1 {
		t.Error("Item inserted is not what came back: ", v)
		return
	}

	expected2 := Task{Script: "Hello"}
	_, err = tasks.Add(&expected2)
	if err != nil {
		t.Error("Failed to add: ", err.Error())
		return
	}
	v = tasks.GetAll()
	if len(v) != 2 {
		t.Error("Expected 2 got ", len(v))
		return
	}
	if *v[0] != expected1 || *v[1] != expected2 {
		t.Error("Items inserted are not what came back: ", v)
		return
	}

	w, _ := tasks.Get(1)
	if *w != expected1 {
		t.Errorf("Tried to get v expected but got %#v", w)
		return
	}
	_, ok := tasks.Get(3)
	if ok {
		t.Error("Tried to get a task that didn't exist but it found something")
		return
	}

	w.Script = "Greetings"
	ok = tasks.Update(w)
	if !ok {
		t.Error("Update failed")
		return
	}
	w2, _ := tasks.Get(1)
	if w != w2 {
		t.Error("Updated value not saved")
		return
	}

	tasks.Delete(1)
	v = tasks.GetAll()
	if len(v) != 1 {
		t.Error("Delete had no effect!")
		return
	}
	w, _ = tasks.Get(2)
	if w.Id != 2 {
		t.Error("Tried to delete task with ID 1 but left with ID ", w.Id)
		return
	}
}
