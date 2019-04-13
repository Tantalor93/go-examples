package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type T struct {
	X string "first"
	Y string "second"
}

type T2 struct {
	X string `json:"x"`
}

func main() {
	t := T{"1", "2"}
	tType := reflect.TypeOf(t)
	fmt.Println(tType)
	fieldX, _ := tType.FieldByName("X")
	fmt.Println(fieldX.Tag)

	t2Type := reflect.TypeOf(T2{})
	field, _ := t2Type.FieldByName("X")
	fmt.Println(field.Tag.Lookup("json"))
	bytes, _ := json.Marshal(T2{"ahoj"})

	fmt.Printf("%s", bytes)
}
