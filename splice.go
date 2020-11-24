// Package splice is a simple implementation of JavaScript's array.splice function for []string in go (golang).
package splice

// Strings emulates JavaScript's array.splice function for string type.
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/splice
//
// start: The index at which to start changing the slice.
// If greater than the length of the slice, start will be set to the length of the slice.
// If negative, it will begin that many elements from the end of the slice.
//
// deleteCount:
// An integer indicating the number of elements in the slice to remove from start.
// It can't be optional in Go, so 0 will act as its absence.
//
// item(s):
// The elements to add to the slice, beginning from start.
// If you do not specify any elements, splice() will only remove elements from the slice.
//
// Differences from JavaScript version:
//
// * Only items can be optional
//
// * One integer argument that trims all from the index can be emulated using (source, index, len(source))
//
// * No support for undefined elements or indices (and strange behavior that comes with them)
//
func Strings(source *[]string, start int, deleteCount int, item ...string) (arrDeletedItems []string) {
	arrDeletedItems = []string{}
	if start > len(*source) {
		start = len(*source)
	}
	if start < 0 {
		start = len(*source) + start
	}
	if start < 0 {
		start = 0
	}
	if deleteCount < 0 {
		deleteCount = 0
	}
	if deleteCount > 0 {
		for i := 0; i < deleteCount; i++ {
			if i+start < len(*source) {
				arrDeletedItems = append(arrDeletedItems, (*source)[i+start])
			}
		}
	}
	deleteCount = len(arrDeletedItems) // Adjust to actual delete count
	grow := len(item) - deleteCount
	switch {
	case grow > 0: // So we grow
		*source = append(*source, make([]string, grow)...)
		copy((*source)[start+deleteCount+grow:], (*source)[start+deleteCount:])
	case grow < 0: // So we shrink
		from := start + len(item)
		to := start + deleteCount
		copy((*source)[from:], (*source)[to:])
		*source = (*source)[:len(*source)+grow]
	}
	copy((*source)[start:], item)
	return
}
