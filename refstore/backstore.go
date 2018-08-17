package refstore

type BackStore struct {
}

func (bs *BackStore) AddId(id Id, e Entry) error {
	return nil
}

func (bs *BackStore) DeleteId(id Id) error {
	return nil
}

func (bs *BackStore) GetEntry(id Id) (Entry, error) {
	return nil, nil
}

func (bs *BackStore) Flush() error {
	return nil
}
