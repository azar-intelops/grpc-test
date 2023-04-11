package test

import (
	"encoding/json"
)

func Add(a int64, b int64) int64 {

	// log.Fatalf("sum of c value is %t shold be ",arguments.a,arguments.b)

	result := a + b
	return result
}

func Subtract(c int64, d int64) int64 {
	res := c - d
	return res
}

func Multiply(c int64, d int64) int64 {
	res := c * d
	return res
}
func Divid(c int64, d int64) int64 {
	res := c / d
	return res
}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}
func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
