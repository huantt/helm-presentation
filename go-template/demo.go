package main

import (
	"os"
	"text/template"
)

func main() {
	type Inventory struct {
		Material string
		Count    uint
	}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}

	// struct input
	sweaters := Inventory{"wool", 17}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	// map input
	err = tmpl.Execute(os.Stdout, map[string]interface{}{
		"Count":    10,
		"Material": "wool",
	})
	if err != nil {
		panic(err)
	}
}
