package test

import (
	"encoding/json"
)

func Add(a int64, b int64) int64 {
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

func IsJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
