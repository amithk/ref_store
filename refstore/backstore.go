package refstore

import "errors"
import "fmt"
import "path/filepath"
import "unsafe"

type StoreMap map[Id]*StoreMapEntry

type BackStore struct {
	spath    string
	storeMap StoreMap
}

type StoreMapEntry struct {
	entry   Entry
	deleted bool
}

type ListEntry struct {
	entry   Entry
	id      Id
	deleted bool
	next    unsafe.Pointer
}

func NewListEntry(e Entry, i Id, d bool) *ListEntry {
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

type FlushCallback func(le *ListEntry) error

func NewBackStore(spath string) *BackStore {
	bs := &BackStore{
		spath:    spath,
		storeMap: make(StoreMap),
	}
	return bs
}

func (bs *BackStore) AddId(id Id, e Entry) error {
	sme := &StoreMapEntry{
		entry:   e,
		deleted: false,
	}
	bs.storeMap[id] = sme
	return nil
}

func (bs *BackStore) DeleteId(id Id) error {
	ent, ok := bs.storeMap[id]
	if !ok {
		return errors.New("Not Found")
	}
	ent.deleted = true
	return nil
}

func (bs *BackStore) GetEntry(id Id) (Entry, error) {
	ent, ok := bs.storeMap[id]
	if !ok {
		return nil, errors.New("Not Found")
	}
	return ent.entry, nil
}

func (bs *BackStore) Flush() error {
	return nil
}

func (bs *BackStore) getFilePathFromId(id Id) string {
	f1 := fmt.Sprintf("%x", id%256)
	f2 := fmt.Sprintf("%x", (id>>8)%256)
	return filepath.Join(bs.spath, f1, f2)
}
