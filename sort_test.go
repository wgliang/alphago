package alphago

import (
	"reflect"
	"testing"
)

type ListInt []int

func (l ListInt) Len() int {
	return len(l)
}

func (l ListInt) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l ListInt) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func Test_InsertSort_Ints(t *testing.T) {
	if !reflect.DeepEqual(InsertSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`InsertSort({2, 8, 5, 1, 7, 4, 3, 9, 6, 0}) != {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func Test_BubbleSort_Ints(t *testing.T) {
	if !reflect.DeepEqual(BubbleSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`BubbleSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}) != ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func Test_BubbleSort_Ints_2(t *testing.T) {
	if !reflect.DeepEqual(BubbleSort(ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`BubbleSort(ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) != ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func Test_SelectSort_Ints(t *testing.T) {
	if !reflect.DeepEqual(SelectSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`SelectSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}) != ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func Test_SelectSort_Ints_2(t *testing.T) {
	if !reflect.DeepEqual(SelectSort(ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`SelectSort(ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) != ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func Test_QuickSort_Ints(t *testing.T) {
	if !reflect.DeepEqual(QuickSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`QuickSort(ListInt{2, 8, 5, 1, 7, 4, 3, 9, 6, 0}) != ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func Test_QuickSort_Ints_2(t *testing.T) {
	if !reflect.DeepEqual(QuickSort(ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}), ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Error(`QuickSort(ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) != ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`)
	}
}

func TestReverse_Ints(t *testing.T) {
	if !reflect.DeepEqual(Reverse(ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}), ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) {
		t.Error(`Reverse(ListInt{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) != ListInt{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}`)
	}
}

type ListString []string

func (l ListString) Len() int {
	return len(l)
}

func (l ListString) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l ListString) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func Test_InsertSort_Strings(t *testing.T) {
	if !reflect.DeepEqual(InsertSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}), ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Error(`InsertSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}) != ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}`)
	}
}

func Test_BubbleSort_Strings(t *testing.T) {
	if !reflect.DeepEqual(BubbleSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}), ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Error(`BubbleSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}) != ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}`)
	}
}

func Test_SelectSort_Strings(t *testing.T) {
	if !reflect.DeepEqual(SelectSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}), ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Error(`SelectSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}) != ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}`)
	}
}

func Test_QuickSort_Strings(t *testing.T) {
	if !reflect.DeepEqual(QuickSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}), ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Error(`QuickSort(ListString{"2", "8", "5", "1", "7", "4", "3", "9", "6", "0"}) != ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}`)
	}
}

func TestReverse_Strings(t *testing.T) {
	if !reflect.DeepEqual(Reverse(ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}), ListString{"9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}) {
		t.Error(`Reverse(ListString{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}) != ListString{"9", "8", "7", "6", "5", "4", "3", "2", "1", "0"}`)
	}
}

type ListStruct struct {
	Name  string
	Age   int
	Score int
}

type ListStructs []ListStruct

func (l ListStructs) Len() int {
	return len(l)
}

func (l ListStructs) Less(i, j int) bool {
	return l[i].Score < l[j].Score
}

func (l ListStructs) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var (
	listStructs = ListStructs{
		ListStruct{
			Name:  "5",
			Age:   20,
			Score: 5,
		},
		ListStruct{
			Name:  "7",
			Age:   20,
			Score: 7,
		},
		ListStruct{
			Name:  "0",
			Age:   20,
			Score: 0,
		},
		ListStruct{
			Name:  "3",
			Age:   20,
			Score: 3,
		},
		ListStruct{
			Name:  "8",
			Age:   20,
			Score: 8,
		},
		ListStruct{
			Name:  "2",
			Age:   20,
			Score: 2,
		},
		ListStruct{
			Name:  "1",
			Age:   20,
			Score: 1,
		},
		ListStruct{
			Name:  "9",
			Age:   20,
			Score: 9,
		},
		ListStruct{
			Name:  "4",
			Age:   20,
			Score: 4,
		},
		ListStruct{
			Name:  "6",
			Age:   20,
			Score: 6,
		},
	}
	listStructsSorted = ListStructs{
		ListStruct{
			Name:  "0",
			Age:   20,
			Score: 0,
		},
		ListStruct{
			Name:  "1",
			Age:   20,
			Score: 1,
		},
		ListStruct{
			Name:  "2",
			Age:   20,
			Score: 2,
		},
		ListStruct{
			Name:  "3",
			Age:   20,
			Score: 3,
		},
		ListStruct{
			Name:  "4",
			Age:   20,
			Score: 4,
		},
		ListStruct{
			Name:  "5",
			Age:   20,
			Score: 5,
		},
		ListStruct{
			Name:  "6",
			Age:   20,
			Score: 6,
		},
		ListStruct{
			Name:  "7",
			Age:   20,
			Score: 7,
		},
		ListStruct{
			Name:  "8",
			Age:   20,
			Score: 8,
		},
		ListStruct{
			Name:  "9",
			Age:   20,
			Score: 9,
		},
	}
	listStructsReverse = ListStructs{
		ListStruct{
			Name:  "9",
			Age:   20,
			Score: 9,
		},
		ListStruct{
			Name:  "8",
			Age:   20,
			Score: 8,
		},
		ListStruct{
			Name:  "7",
			Age:   20,
			Score: 7,
		},
		ListStruct{
			Name:  "6",
			Age:   20,
			Score: 6,
		},
		ListStruct{
			Name:  "5",
			Age:   20,
			Score: 5,
		},
		ListStruct{
			Name:  "4",
			Age:   20,
			Score: 4,
		},
		ListStruct{
			Name:  "3",
			Age:   20,
			Score: 3,
		},
		ListStruct{
			Name:  "2",
			Age:   20,
			Score: 2,
		},
		ListStruct{
			Name:  "1",
			Age:   20,
			Score: 1,
		},
		ListStruct{
			Name:  "0",
			Age:   20,
			Score: 0,
		},
	}
)

func Test_InsertSort_Structs(t *testing.T) {
	if !reflect.DeepEqual(InsertSort(listStructs), listStructsSorted) {
		t.Error("InsertSort error")
	}
}

func Test_BubbleSort_Structs(t *testing.T) {
	if !reflect.DeepEqual(BubbleSort(listStructs), listStructsSorted) {
		t.Error("BubbleSort error")
	}
}

func Test_SelectSort_Structs(t *testing.T) {
	if !reflect.DeepEqual(SelectSort(listStructs), listStructsSorted) {
		t.Error("SelectSort error")
	}
}

func Test_QuickSort_Structs(t *testing.T) {
	if !reflect.DeepEqual(QuickSort(listStructs), listStructsSorted) {
		t.Error("QuickSort error")
	}
}

func Test_Reverse_Structs(t *testing.T) {
	if !reflect.DeepEqual(Reverse(listStructsSorted), listStructsReverse) {
		t.Error("Reverse error")
	}
}
