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

type AnimalExtend struct {
	ID    string
	Title string
	Name  string `reorm:"index"`
	Color string
	Type  int
	Sound Sound
}

type Sound struct {
	Message string
}

type WithoutDefault struct {
	Message string
	Name    string `reorm:"primaryKey"`
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

	c3 := &Car{
		ID:    3,
		Name:  "Mercedes1",
		Color: "Black2",
		Attributes: Attributes{
			Windows: 4,
			Smart:   true,
		},
	}
	r, err := reorm.New(&reorm.Config{Address: "localhost:6379"})
	fmt.Println(r.Save(c))
	fmt.Println(r.Save(c2))
	r.Save(c3)
	var resp Car
	if err := r.GetByID("Car", 1, &resp); err != nil {
		panic(err)
	}
	fmt.Println("Get BY ID 1: ", resp)

	var resp2 Car
	if err := r.Get(&Car{ID: 2}, &resp2); err != nil {
		panic(err)
	}
	fmt.Println("RESP ID2: ", resp2)

	var resp3 Car
	if err := r.Get(&Car{Name: "Mercedes1"}, &resp3); err != nil {
		panic(err)
	}
	fmt.Println("RESP3: ", resp3)

	var resp4 Car
	if err := r.Last(&Car{Name: "Mercedes1"}, &resp4); err != nil {
		panic(err)
	}
	fmt.Println("RESP4: ", resp4)

	resInter, _ := r.Find(&Car{
		Color: "Black1",
	})
	for _, d := range resInter {
		fmt.Println(d.(*Car))
	}

	if err := r.Update(&Car{ID: 2, Name: "Mercedes1"}, &Car{Name: "Mercedes20"}); err != nil {
		panic(err)
	}

	var resp6 Car
	if err := r.Get(&Car{ID: 2}, &resp6); err != nil {
		panic(err)
	}
	fmt.Println("RESP6: ", resp6)

	a2 := &AnimalExtend{
		Title: "Dog",
		Name:  "Bob",
		Color: "Black",
		Type:  1,
		Sound: Sound{
			Message: "Data",
		},
	}
	err = r.Save(a2)
	if err != nil {
		panic(err)
	}

	var resp7 AnimalExtend
	if err := r.Get(&AnimalExtend{Name: "Bob"}, &resp7); err != nil {
		panic(err)
	}
	fmt.Println("RESP6: ", resp7)

	resp8, err := r.GetValueByField("Car", "Name", &Car{ID: 1})
	if err != nil {
		panic(err)
	}

	fmt.Println("ATTR: ", resp8.(string))

}
