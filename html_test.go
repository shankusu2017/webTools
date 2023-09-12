package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	var testSlice []statement

	var item, item2 statement
	item.Weight = 30
	item.Text = "hi30"
	item2.Weight = 20
	item2.Text = "hi20"

	testSlice = append(testSlice, item)
	testSlice = append(testSlice, item2)
	b, _ := json.Marshal(testSlice)
	fmt.Printf("[%s]\n", string(b))
}
