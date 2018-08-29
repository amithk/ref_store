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

	err = ll.Flush(testFlush, testFlushEntry)
	if err != nil {
		fmt.Println("Error", err, "In Flush")
	}
}

func testFlushEntry(le *ListEntry) error {
	fmt.Println("testFlushEntryCalled")
	le.Print()
	return nil
}

func testFlush() error {
	fmt.Println("testFlush called")
	return nil
}
