package util

import "fmt"

func GenericArray(arr []int) []interface{} {
	t := make([]interface{}, len(arr))
	for i, v := range arr {
		t[i] = v
	}
	return t
}

func IntType(t interface{}) int {
	x, ok := t.(int)
	if !ok {
		panic("It is not an integer")
	}
	return x
}

func CompareInt(a interface{}, b interface{}) int {
	if a == nil && b == nil {
		return 0
	}

	if a == nil {
		return 1
	} else if b == nil {
		return -1
	}

	v1, ok1 := a.(int)
	if !ok1 {
		panic(fmt.Sprintf("unable to convert %T as int", a))
	}
	v2, ok2 := b.(int)
	if !ok2 {
		panic(fmt.Sprintf("unable to convert %T as int", b))
	}
	return v2 - v1
}
