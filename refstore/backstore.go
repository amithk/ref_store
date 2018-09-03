package refstore

import "fmt"
import "path/filepath"
import "sync"
import "unsafe"

// TODO: Restructure list entry. Let linked list store the next pointer.
// This ensures all the unsafe pointer logic to be moved to linked list
// implementation. Note: this may lead to more allocations.
type ListEntry struct {
	entry   Entry
	id      Id
	deleted bool
	next    unsafe.Pointer
}

func NewListEntry(e Entry, i Id, d bool) *ListEntry {
	// TODO: sync pool can be used here. Custom slab allocator can be used.
	return &ListEntry{
		entry:   e,
		id:      i,
		deleted: d,
	}
}

func (le *ListEntry) Print() {
	if le == nil {
		fmt.Println("nil")
		return
	}

	fmt.Printf("id: %v, entry: %v, deleted = %v\n", le.id, le.entry, le.deleted)
}

type FlushCallback func() error
type FlushEntryCallback func(le *ListEntry) error

type FlushMap map[Id]Entry

type BackStore struct {
	spath  string
	ll     *LinkedList
	addMap FlushMap
	delMap FlushMap    // TODO: use better data type for delMap
	fLock  *sync.Mutex // Allow only one flush operation at a time
}

func NewBackStore(spath string) *BackStore {
	bs := &BackStore{
		spath:  spath,
		ll:     NewLinkedList(),
		addMap: make(FlushMap),
		delMap: make(FlushMap),
		fLock:  new(sync.Mutex),
	}
	return bs
}

func (bs *BackStore) AddId(id Id, e Entry) error {
	le := NewListEntry(e, id, false)
	err := bs.ll.Append(le)
	return err
}

func (bs *BackStore) DeleteId(id Id) error {
	le := NewListEntry(nil, id, true)
	err := bs.ll.Append(le)
	return err
}

func (bs *BackStore) GetEntry(id Id) (Entry, error) {
	// TODO: Implement read logic
	// Need to implement read caching
	return nil, nil
}

func (bs *BackStore) Flush() error {
	bs.fLock.Lock()
	defer bs.fLock.Unlock()

	err := bs.ll.Flush(bs.flushAll, bs.flushEntry)
	return err
}

func (bs *BackStore) getFilePathFromId(id Id) string {
	f1 := fmt.Sprintf("%x", id%256)
	f2 := fmt.Sprintf("%x", (id>>8)%256)
	return filepath.Join(bs.spath, f1, f2)
}

func (bs *BackStore) flushEntry(le *ListEntry) error {
	// TODO: using golang map might not be perf efficient
	// Need to reconsider if need arises.
	if !le.deleted {
		bs.addMap[le.id] = le.entry
	} else {
		if _, ok := bs.addMap[le.id]; ok {
			delete(bs.addMap, le.id)
		} else {
			bs.delMap[le.id] = nil
		}
	}
	return nil
}

func (bs *BackStore) flushAll() error {
	// TODO: Implement actual flush logic.
	fmt.Println("bs.addMap =", bs.addMap)
	fmt.Println("bs.delMap =", bs.delMap)
	return nil
}
