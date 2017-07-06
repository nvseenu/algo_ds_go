package sort

type SortName int

const (
	BUBBLE SortName = iota
	QUICK
	SELECTION
	INSERTION
)

func NewSorter(name SortName, compare func(a interface{}, b interface{}) int) Sorter {

	var sorter Sorter = nil

	switch name {
	case BUBBLE:
		sorter = typeSwitch(&BubbleSort{compare})
	case QUICK:
		sorter = typeSwitch(&QuickSort{compare})
	case SELECTION:
		sorter = typeSwitch(&SelectionSort{compare})
	case INSERTION:
		sorter = typeSwitch(&InsertionSort{compare})
	}
	return sorter
}

func typeSwitch(s interface{}) Sorter {
	return s.(Sorter)
}
