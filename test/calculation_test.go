package test

import (
	"testing"
)

type Arguments struct {
	a int64
	b int64
	c int64
}

func TestAdd(t *testing.T) {
	data := []Arguments{{10, 20, 30}, {2, 3, 5}}

	for _, val := range data {
		cal_add := Add(val.a, val.b)
		if cal_add != val.c {
			t.Errorf("got %v but got %v", val.c, cal_add)
		}
	}

}
func TestSubtract(t *testing.T) {
	data := Arguments{30, 20, 10}
	cal_add := Subtract(data.a, data.b)

	if cal_add != data.c {
		t.Errorf("got %v but got %v", data.c, cal_add)
	}

}
func TestMultiply(t *testing.T) {

	data := Arguments{3, 6, 18}
	cal_add := Multiply(data.a, data.b)

	if cal_add != data.c {
		t.Errorf("got %v but need %v", data.c, cal_add)
	}

}

func TestDivid(t *testing.T) {
	// var data.c int64 = 10
	data := Arguments{12, 6, 2}
	cal_add := Divid(data.a, data.b)

	if cal_add != data.c {
		t.Errorf("got %v but need %v", data.c, cal_add)
	}

}
