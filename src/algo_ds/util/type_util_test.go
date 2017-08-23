package util

import "testing"

type compareData struct {
	a   interface{}
	b   interface{}
	res int
}

var compareDataSet = []compareData{
	{5, 3, -2},
	{3, 5, 2},
	{3, 3, 0},
	{nil, nil, 0},
	{nil, 3, 1},
	{5, nil, -1},
}

func TestCompareInt(t *testing.T) {
	for _, data := range compareDataSet {
		res := CompareInt(data.a, data.b)
		if res != data.res {
			t.Errorf("Expected %v but got %v for data %v", data.res, res, data)
		}
	}

}
