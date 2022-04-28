package utils

import (
	"encoding/json"
	"fmt"
)

// Dump
// @Description: dump data
// @param value
// @return string
func Dump(value any) {
	response, e := json.MarshalIndent(value, "", "  ")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(response))
}
