package main

import (
	"fmt"

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
	fmt.Println(r.Save(c))

	var resp Car
	if err := r.GetByID(1, &resp); err != nil {
		panic(err)
	}
	fmt.Println("RESP: ", resp)
	if err := r.DeleteByID(1); err != nil {
		panic(err)
	}
	if err := r.GetByID(1, &resp); err != nil {
		panic(err)
	}
	fmt.Println("RESP: ", resp)
}
