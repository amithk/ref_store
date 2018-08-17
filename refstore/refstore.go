package refstore

type Entry []byte

type Id uint64

type Reference []byte

// Implements RefStoreInterface
type RefStore struct {
}

func (rs *RefStore) AddEntryReference(e Entry, r Reference) (Id, error) {
	return 0, nil
}

func (rs *RefStore) DeleteEntryReference(e Entry, r Reference) error {
	return nil
}

func (rs *RefStore) DeleteId(id Id) error {
	return nil
}

func (rs *RefStore) GetIdFromEntry(e Entry) (Id, error) {
	return 0, nil
}

func (rs *RefStore) GetEntryFromId(id Id) (Entry, error) {
	return nil, nil
}

func (rs *RefStore) AddReferenceById(id Id, r Reference) error {
	return nil
}

func (rs *RefStore) RemoveReferenceById(id Id, r Reference) error {
	return nil
}

func (rs *RefStore) Checkpoint() error {
	return nil
}
