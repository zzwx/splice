// Package splice is an implementation of JavaScript's array.splice function for []T.
package splice

// Splice modifies the source slice by deleting delete amount of items starting with index and replacing them with
// optional item(s). It returns all the elements that have been removed, nil for zero removed items.
//
// Splice emulates JavaScript's array.splice function for type T
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/splice
//
// start: The index at which to start changing the slice.
//
// If greater than the length of the slice, start will be set to the length of the slice.
// If negative, it will begin that many elements from the end of the slice.
//
// delete:
//
// An integer indicating the number of elements in the slice to remove from start.
// It can't be optional in Go, so 0 will act as its absence.
//
// item(s):
//
// The elements to add to the slice, beginning from start.
// If you do not specify any elements, splice() will only remove elements from the slice.
//
// Differences from the JavaScript version:
//
// * Only item(s) arguments may be omitted. The one integer argument call that only trims the source from start
// to the end of the source can be emulated using Splice(source, index, len(source)).
//
// * No support for undefined elements or indices.
func Splice[T any](source *[]T, start int, delete int, item ...T) (removed []T) {
	if start > len(*source) {
		start = len(*source)
	}
	if start < 0 {
		start = len(*source) + start
	}
	if start < 0 {
		start = 0
	}
	if delete < 0 {
		delete = 0
	}
	if delete > 0 {
		for i := 0; i < delete; i++ {
			if i+start < len(*source) {
				removed = append(removed, (*source)[i+start])
			}
		}
	}
	delete = len(removed) // Adjust to actual delete count
	grow := len(item) - delete
	switch {
	case grow > 0: // So we grow
		*source = append(*source, make([]T, grow)...)
		copy((*source)[start+delete+grow:], (*source)[start+delete:])
	case grow < 0: // So we shrink
		from := start + len(item)
		to := start + delete
		copy((*source)[from:], (*source)[to:])
		*source = (*source)[:len(*source)+grow]
	}
	copy((*source)[start:], item)
	return
}
