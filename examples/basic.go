package main

import (
	reorm "github.com/saromanov/re-orm"
)

type Car struct {
	ID    int64
	Name  string
	Color string
}

func main() {
	c := &Car{
		ID:    1,
		Name:  "BMW",
		Color: "Black",
	}
	r := reorm.New(&reorm.Config{})
	r.Save(c)
}
