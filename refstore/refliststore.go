package refstore

type RefListStore struct {
}

func (rls *RefListStore) AddRef(id Id, ref Reference) error {
	return nil
}

func (rls *RefListStore) DeleteRef(id Id, ref Reference) error {
	return nil
}

func (rls *RefListStore) PurgeAllRefs(id Id) error {
	return nil
}

func (rls *RefListStore) Flush() error {
	return nil
}
