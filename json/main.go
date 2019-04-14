package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Object struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func main() {
	s := `{"name":"ahoj", "number": 1 }`

	var res map[string]interface{}
	json.Unmarshal([]byte(s), &res)
	fmt.Println(res)

	var res2 map[string]interface{}
	sReader := strings.NewReader(s)
	json.NewDecoder(sReader).Decode(&res2)
	fmt.Println(res2)

	var res3 Object
	json.Unmarshal([]byte(s), &res3)
	fmt.Println(res3)

	bytes, _ := json.Marshal(res3)
	fmt.Println(string(bytes))
}
