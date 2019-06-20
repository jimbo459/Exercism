package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

//IntList ...
type IntList struct {
	list *[]int
}

func (l *IntList) foldl(list *[]int, f binFunc) []int {
	tmpList := *list
	return tmpList
}

// TODO
// foldl
// foldr
// filter - returns a new array with all items which pass the specified criteria (function called on every item in the array.)
// length
// map - calls a method on every item in the array.
// reverse
// append
// concat
// Folding is the same as reduce in Javascript -
// - Take an identity (accumulator, int/bool)
// - Take an operation to perform on the identity / accumulated value
// - And a list on which to iterate. return fold(false, or, list);

// Create a function which takes in a list and a function, then applies the function to each element in the list.
