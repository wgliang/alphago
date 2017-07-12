package alphago

// Vector is a struct that as a parametric type abstraction of
// Generic method. It is a interface type. And the user should
// implement function:Len(), Less() and Swap().
type Vector interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// InsertSort is a O(n^2) time complexity sorting algorithm. And
// it adapt to an already ordered or partially ordered data sequence.
func InsertSort(vector Vector) Vector {
	l := vector.Len() - 1
	for i := 1; i <= l; i++ {
		for j := i; j > 0 && vector.Less(j, j-1); j-- {
			vector.Swap(j, j-1)
		}
	}
	return vector
}

// BubbleSort is a O(n^2) time complexity sorting algorithm. It will
// repeatedly visit the sequence to be sorted, compare two elements at
// a time, if they are wrong in order to exchange them over.
func BubbleSort(vector Vector) Vector {
	l := vector.Len() - 1
	for i := 0; i < l; i++ {
		for j := l; j > i; j-- {
			if vector.Less(j, j-1) {
				vector.Swap(j, j-1)
			}
		}
	}
	return vector
}

// SelectSort is a O(n^2) time complexity sorting algorithm. It's running
// time and input has nothing to do, the smallest data movement, when the
// amount of data is relatively small when applicable.
func SelectSort(vector Vector) Vector {
	l := vector.Len() - 1
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j <= l; j++ {
			if vector.Less(j, min) {
				min = j
			}
		}
		vector.Swap(i, min)
	}
	return vector
}

// QuickSort is a O(nlogn) time complexity sorting algorithm. And it is the
// fastest universal sorting algorithm, and in most real cases, quick sorting
// is the best choice.
func QuickSort(vector Vector) Vector {
	quicksort(vector, 0, vector.Len()-1)
	return vector
}

func quicksort(vector Vector, l int, r int) {
	if l >= r {
		return
	}

	mid := l
	i := l + 1

	for j := l; j <= r; j++ {
		if vector.Less(j, mid) {
			vector.Swap(i, j)
			i++
		}
	}

	vector.Swap(l, i-1)

	quicksort(vector, l, i-2)
	quicksort(vector, i, r)
}

// Reverse is a function that will reverse vector.
func Reverse(vector Vector) Vector {
	l := vector.Len() - 1
	r := vector.Len() / 2
	for i := 0; i < r; i++ {
		vector.Swap(i, l-i)
	}
	return vector
}
