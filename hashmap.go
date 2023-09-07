package main

const TableSize = 8

// Hash map implementation using chaining method of collision resolution.
type HashMap struct {
	table [TableSize]*Bucket
}

func Init() *HashMap {
	result := &HashMap{}
	for i := range result.table {
		result.table[i] = &Bucket{}
	}
	return result
}

func (h *HashMap) getIndex(key uint32) uint32 {
	result := Digest(key, TableSize)
	println("[DBG] index is ", result)
	return result
}

// Check whether the map contains any data for the requested key.
func (h *HashMap) Has(key uint32) bool {
	index := h.getIndex(key)
	found := h.table[index].Search(key)
	return found != nil
}

func (h *HashMap) Get(key uint32) any {
	index := h.getIndex(key)
	if found := h.table[index].Search(key); found != nil {
		return found.Value
	}
	return nil
}

// Returns true if inserted or false if the map contents remained the same.
func (h *HashMap) Insert(key uint32, value string) bool {
	if h.Has(key) {
		return false

	} else {
		index := h.getIndex(key)
		h.table[index].Append(&MapItem{
			Key:   key,
			Value: value,
		})
		return true
	}
}

func (h *HashMap) Delete(key uint32) {
	// TODO: implement
	// index := h.getIndex(key)
}
