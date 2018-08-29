package refstore

import "fmt"
import "testing"

func TestBackStoreBasic(t *testing.T) {
	bs := NewBackStore("")

	bs.AddId(10, []byte("aaa"))
	bs.AddId(11, []byte("bbb"))
	bs.AddId(12, []byte("ccc"))
	bs.AddId(13, []byte("ddd"))
	bs.DeleteId(11)
	bs.DeleteId(22)

	err := bs.Flush()
	if err != nil {
		fmt.Println("Error in flush", err)
	}
}
