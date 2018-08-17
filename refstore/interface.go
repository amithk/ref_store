/*

  Terminology:
  Entry: the byte strem to be deduplicated
  Id:	 the unique identifier to be used in place of the entry

*/

package refstore

type RefStoreInterface interface {

	// Add Entry and Reference to the store. If the entry is already
	// present, only reference is added to the store.
	AddEntryReference(Entry, Reference) (Id, error)

	// Delete a reference. Mark entry as deleted only if this is the
	// last reference being deleted.
	DeleteEntryReference(Entry, Reference) error

	// Mark Id and Corresponding Entry as deleted. Remove all
	// references by force.
	DeleteId(Id) error

	GetIdFromEntry(Entry) (Id, error)

	GetEntryFromId(Id) (Entry, error)

	AddReferenceById(Id, Reference) error

	RemoveReferenceById(Id, Reference) error

	// All the writes happen in memory. Background thread ensures
	// eventual data flush, but application can call this function
	// to force the updates to be written to the disk.
	Checkpoint() error
}
