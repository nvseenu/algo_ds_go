package util

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
