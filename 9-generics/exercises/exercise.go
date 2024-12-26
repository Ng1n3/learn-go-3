package exercises

func Exercise1() {
	// Write a function called "Filter" that takes a slice of any type and a function that takes an element of that type and returns a bool.
	// Filter should return a new slice containing only the elements for which the function returns true.
	// For example, calling Filter([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }) should return []int{2, 4}.
}

func Filter[T any](slice []T, sliceFunc func(T) bool) []T {
	var res []T
	for _, v := range slice {
		if sliceFunc(v) {
			res = append(res, v)
		}
	}
	return res
}

func Exercise2() {
	// Write a function called "Map" that takes a slice of any type and a function that takes an element of that type and returns a new element of a different type.
	// Map should return a new slice containing the new elements.
	// For example, calling Map([]int{1, 2, 3, 4}, func(i int) string { return fmt.Sprintf("%d", i) }) should return []string{"1", "2", "3", "4"}.
}

func Map[T any, U any](slice []T, mapper func(T) U) []U {
	var res []U
	for _, v := range slice {
		res = append(res, mapper(v))
	}
	return res
}

func Exercise3() {
	// Write a function called "Reduce" that takes a slice of any type and a function that takes two elements of that type and returns a new element of that type.
	// Reduce should return a single element of that type.
	// For example, calling Reduce([]int{1, 2, 3, 4}, func(i, j int) int { return i + j }) should return 10.
}

func Reduce[T any](slice []T, reducer func(T, T) T) T {
	res := slice[0]
	for _, v := range slice[1:] {
		res = reducer(res, v)
	}
	return res
}

func Exercise4() {
	// Write a function called "All" that takes a slice of any type and a function that takes an element of that type and returns a bool.
	// All should return true if the function returns true for all elements in the slice, and false otherwise.
	// For example, calling All([]int{1, 2, 3, 4}, func(i int) bool { return i < 5 }) should return true.
}

func All[T any](slice []T, sliceFunc func(T) bool) bool {
	for _, v := range slice {
		if sliceFunc(v) == false {
			return false
		}
	}
	return true
}

func Exercise5() {
	// Write a function called "Any" that takes a slice of any type and a function that takes an element of that type and returns a bool.
	// Any should return true if the function returns true for any element in the slice, and false otherwise.
	// For example, calling Any([]int{1, 2, 3, 4}, func(i int) bool { return i > 3 }) should return true.
}

func Any[T any](slice []T, sliceFunc func(T) bool) bool {
	for _, v := range slice {
		if sliceFunc(v) {
			return true
		}
	}
	return false
}

func Exercise6() {
  // Write a function called "Find" that takes a slice of any type and a function that takes an element of that type and returns a bool.
  // Find should return the first element in the slice for which the function returns true, or the zero value of the type if no such element is found.
  // For example, calling Find([]int{1, 2, 3, 4}, func(i int) bool { return i > 2 }) should return 3.
}

func Find[T any](slice []T, sliceFunc func(T) bool) T {
  for _, v := range slice {
    if sliceFunc(v) {
      return v
    }
  }
  var zero T
  return zero
}
