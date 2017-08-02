package collections

type collError struct {
	message string
}

func (e *collError) Error() string {
	return "CollError => " + e.message
}
