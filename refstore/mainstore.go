package refstore

type MainStore struct {
	spath string
}

func NewMainStore(spath string) *MainStore {
	ms := &MainStore{
		spath: spath,
	}
	return ms
}

func (ms *MainStore) AddEntry(e Entry, id Id) error {
	return nil
}

func (ms *MainStore) DeleteEntry(e Entry) error {
	return nil
}

func (ms *MainStore) GetId(e Entry) (Id, error) {
	return 0, nil
}

func (ms *MainStore) Flush() error {
	return nil
}
