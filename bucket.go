package main

type entry struct {
	mapItem *MapItem
	next    *entry
}

// A linked list to represent entries with the same hash.
type Bucket struct {
	head *entry
}

func (b *Bucket) Append(mapItem *MapItem) {
	if b.head == nil {
		b.head = &entry{
			mapItem: mapItem,
		}
		return
	}

	cur := b.head
	for cur.next != nil {
		cur = cur.next
	}

	newEntry := &entry{
		mapItem: mapItem,
	}
	cur.next = newEntry
}

func (b *Bucket) Search(key uint32) *MapItem {
	if b.head == nil {
		return nil
	}

	cur := b.head
	for cur != nil {
		if cur.mapItem.Key == key {
			return cur.mapItem
		}
		cur = cur.next
	}

	return nil
}
