package splice

import (
	"fmt"
	"testing"
)

func ExampleStrings() {
	source := []string{"a"}
	deleted := Strings(&source, 0, 1, "0", "1")
	fmt.Println("source:", source)
	fmt.Println("deleted:", deleted)
	// Emulate one argument source.slice(1)
	source = []string{"a", "b", "c"}
	deleted = Strings(&source, 1, len(source))
	fmt.Println("source:", source)
	fmt.Println("deleted:", deleted)
	// Output:
	// source: [0 1]
	// deleted: [a]
	// source: [a]
	// deleted: [b c]
}

func TestSpliceString(t *testing.T) {
	var source, deleted []string

	for i := -1000; i < 1000; i++ {
		source = []string{}
		deleted = Strings(&source, i, 100-i, "0", "1")
		expect(t, "s", source, "0", "1")
		expect(t, "d", deleted)
	}

	source = []string{"a"}
	deleted = Strings(&source, 0, 1, "0", "1")
	expect(t, "s", source, "0", "1")
	expect(t, "d", deleted, "a")

	source = []string{"a"}
	deleted = Strings(&source, 2, 1, "0", "1")
	expect(t, "s", source, "a", "0", "1")
	expect(t, "d", deleted)

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 2, 2, "0", "1", "2", "3", "4")
	expect(t, "s", source, "a", "b", "0", "1", "2", "3", "4", "e")
	expect(t, "d", deleted, "c", "d")

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 2, 0, "0", "1", "2", "3", "4")
	expect(t, "s", source, "a", "b", "0", "1", "2", "3", "4", "c", "d", "e")
	expect(t, "d", deleted)

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 0, 2)
	expect(t, "s", source, "c", "d", "e")
	expect(t, "d", deleted, "a", "b")

	// No-op: no delete count, no items to insert
	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 100, 0)
	expect(t, "s", source, "a", "b", "c", "d", "e")
	expect(t, "d", deleted)

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 100, 5, "0", "1", "2", "3", "4")
	expect(t, "s", source, "a", "b", "c", "d", "e", "0", "1", "2", "3", "4")
	expect(t, "d", deleted)

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 2, 2, "0")
	expect(t, "s", source, "a", "b", "0", "e")
	expect(t, "d", deleted, "c", "d")

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 2, 3, "0")
	expect(t, "s", source, "a", "b", "0")
	expect(t, "d", deleted, "c", "d", "e")

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, 0, 3, "0", "1")
	expect(t, "s", source, "0", "1", "d", "e")
	expect(t, "d", deleted, "a", "b", "c")

	source = []string{"a"}
	deleted = Strings(&source, 0, 3, "0", "1")
	expect(t, "s", source, "0", "1")
	expect(t, "d", deleted, "a")

	source = []string{"a"}
	deleted = Strings(&source, 0, 0, "0", "1")
	expect(t, "s", source, "0", "1", "a")
	expect(t, "d", deleted)

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, -5, 2, "0", "1", "2", "3", "4")
	expect(t, "s", source, "0", "1", "2", "3", "4", "c", "d", "e")
	expect(t, "d", deleted, "a", "b")

	source = []string{"a", "b", "c", "d", "e"}
	deleted = Strings(&source, -5, 2)
	expect(t, "s", source, "c", "d", "e")
	expect(t, "d", deleted, "a", "b")

	source = []string{"0", "3"}
	deleted = Strings(&source, 1, 0, "1", "2")
	expect(t, "s", source, "0", "1", "2", "3")
	expect(t, "d", deleted)

	source = []string{"0", "3"}
	deleted = Strings(&source, 1, 1, "1", "2")
	expect(t, "s", source, "0", "1", "2")
	expect(t, "d", deleted, "3")

	source = []string{"0", "3"}
	deleted = Strings(&source, 1, 10, "1", "2")
	expect(t, "s", source, "0", "1", "2")
	expect(t, "d", deleted, "3")

	source = []string{"0", "3"}
	deleted = Strings(&source, -1, -1, "1", "2")
	expect(t, "s", source, "0", "1", "2", "3")
	expect(t, "d", deleted)

	source = []string{"0", "3"}
	deleted = Strings(&source, 0, -10, "1", "2")
	expect(t, "s", source, "1", "2", "0", "3")
	expect(t, "d", deleted)

	source = []string{"", "1", "", "3"}
	deleted = Strings(&source, 1, 2, "1", "2")
	expect(t, "s", source, "", "1", "2", "3")
	expect(t, "d", deleted, "1", "")

	source = []string{"0", "1", "2"}
	deleted = Strings(&source, 1, 0, "?")
	expect(t, "s", source, "0", "?", "1", "2")
	expect(t, "d", deleted)
}

func expect(t *testing.T, wh string, src []string, items ...string) {
	if len(src) != len(items) {
		goto no
	}
	for i := range src {
		if src[i] != items[i] {
			goto no
		}
	}
	return
no:
	t.Errorf("%v: got: %+v, want: %+v\n", wh, src, items)
}
