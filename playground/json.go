package main

import (
	"encoding/json"
	"fmt"
)

type human struct {
	Name string `json:"name"`
}

func main() {
	h := human{"John"}
	b, _ := json.Marshal(h)
	fmt.Printf("%v\n", string(b))
}
