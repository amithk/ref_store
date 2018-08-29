package refstore

import "sync/atomic"
import "unsafe"

type LinkedList struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLinkedList() *LinkedList {
	l := &LinkedList{}
	return l
}

func (l *LinkedList) Append(le *ListEntry) error {
	var swapped bool
	ule := unsafe.Pointer(le)
	for {
		t := atomic.LoadPointer(&l.tail)
		if t == nil {
			// Safely assume that head too is nil
			swapped = atomic.CompareAndSwapPointer(&l.tail, nil, ule)
			if !swapped {
				continue
			}

			swapped = atomic.CompareAndSwapPointer(&l.head, nil, ule)
			if !swapped {
				continue
			}
		} else {
			let := (*ListEntry)(t)
			swapped = atomic.CompareAndSwapPointer(&let.next, nil, ule)
			if !swapped {
				continue
			}

			swapped = atomic.CompareAndSwapPointer(&l.tail, t, ule)
			if !swapped {
			}
		}
		break
	}
	// ideally there should be timeout for this.
	return nil
}

func (l *LinkedList) Flush(flush FlushCallback, flushentry FlushEntryCallback) error {
	var err error

	t := atomic.LoadPointer(&l.tail)
	if t == nil {
		return nil
	}

	h := atomic.LoadPointer(&l.head)
	if h == nil {
		return nil
	}

	for {
		err = flushentry((*ListEntry)(h))
		if err != nil {
			return err
		}

		leh := (*ListEntry)(h)
		h = leh.next

		if h == t {
			err = flushentry((*ListEntry)(h))
			if err != nil {
				return err
			}
			break
		}
	}

	err = flush()
	if err != nil {
		return err
	}

	let := (*ListEntry)(t)
	atomic.StorePointer(&l.head, let.next)
	return nil
}
