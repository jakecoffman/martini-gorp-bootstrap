package main

import (
	"html/template"

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
