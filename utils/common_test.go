package utils

import "testing"

func Test_IsEmpty(t *testing.T) {
	exp := true
	res := IsEmpty("")
	if exp != res {
		t.Fatalf("got %t, wanted %t", res, exp)
	}
}
