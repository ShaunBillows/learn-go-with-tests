Arrays:

- Size is encoded in type (e.g. type of [4]int != type of [5]int).
- To create an array: myArray := [3]int{1, 2, 3}.

Slices:

- Do not encode type (dynamic data structure).
- To create a slice from an array: mySlice := myArray[1:3].
- To create a slice: mySlice := []int{1, 2, 3}.

reflect.DeepEqual:

- Allows you to check equality with slices.
- Usage: reflect.DeepEqual(want, got).
- Warning: Not type safe (e.g. will compile even if args are of different types).

Coverage tool:

- Built-in tool.
- Helps identify areas of your code not covered by tests.
- Usage: 'go test -cover'.

variadic functions:

- Takes a variable number of arguments.
- Usage: 'func myFunc(args ...[]int) []int {...}'.

make:

- Used to create a new slice, map, or channel with a specified length and capacity.
- Usage: a := make([]int, 5) // Returns slice []int{0, 0, 0, 0, 0} with len(a)=5 and cap(a)=5.
- Initialised with zero values for its elements.

append:

- Go language feature.
- Adds an element to the end of a slice.
- Does not modify existing slice. Creates a new slice that points to the same array. Modifies the underlying array only if there wasn't enough capacity to create the slice.

slicing slices:

- Slice[low:high] // Creates new slice.
- If one value is omitted, defaults to minimum/maximum value.
