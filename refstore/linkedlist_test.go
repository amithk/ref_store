package refstore

import "testing"
import "fmt"

func TestLinkedListBasic(t *testing.T) {
	var err error
	ll := NewLinkedList()

	for i := 0; i < 10; i++ {
		e := make(Entry, 10)
		ii := Id(i)
		le := NewListEntry(e, ii, false)
		err = ll.Append(le)
		if err != nil {
			fmt.Println("Error", err, "In Append")
		}
	}

	err = ll.Flush(testFlushCallback)
	if err != nil {
		fmt.Println("Error", err, "In Flush")
	}
}

func testFlushCallback(le *ListEntry) error {
	le.Print()
	return nil
}
