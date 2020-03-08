package main

import (
	"fmt"

	reorm "github.com/saromanov/re-orm"
)

type Car struct {
	ID         int64
	Name       string `reorm:"index"`
	Color      string `reorm:"index"`
	Attributes Attributes
}

type Attributes struct {
	Windows int64
	Smart   bool
}

func main() {
	c := &Car{
		ID:    1,
		Name:  "BMW1",
		Color: "Black1",
		Attributes: Attributes{
			Windows: 4,
			Smart:   true,
		},
	}

	c2 := &Car{
		ID:    2,
		Name:  "Mercedes1",
		Color: "Black1",
		Attributes: Attributes{
			Windows: 3,
			Smart:   true,
		},
	}
	r := reorm.New(&reorm.Config{})
	fmt.Println(r.Save(c))
	fmt.Println(r.Save(c2))
	var resp Car
	if err := r.GetByID("Car", 1, &resp); err != nil {
		panic(err)
	}
	fmt.Println("RESP: ", resp)

	var resp2 Car
	if err := r.Get(&Car{ID: 2}, &resp2); err != nil {
		panic(err)
	}
	fmt.Println("RESP2: ", resp2)
	/*if err := r.DeleteByID(1); err != nil {
		panic(err)
	}
	if err := r.GetByID(1, &resp); err != nil {
		panic(err)
	}
	fmt.Println("RESP: ", resp)*/
	var respAll []Car
	r.Find(&Car{
		Color: "Black1",
	}, &respAll)
	fmt.Println("RERR: ", respAll)
}
