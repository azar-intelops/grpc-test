package utils

import (
	"fmt"
	"testing"
)

type IsEmpty_TestCases struct {
	value    string
	expected bool
}

func Test_IsEmpty(t *testing.T) {
	data := []IsEmpty_TestCases{
		{
			value:    "",
			expected: true,
		},
		{
			value:    "world",
			expected: false,
		},
		{
			value:    "hello",
			expected: true,
		},
	}
	for _, testcase := range data {
		t.Run(fmt.Sprintf("IsEmpty(%v) == %t", testcase.value, testcase.expected), func(t *testing.T) {
			res := IsEmpty(testcase.value)
			if testcase.expected != res {
				t.Fatalf("got %t, wanted %t", res, testcase.expected)
			}
		})

	}
}
