package main

import (
	"fmt"
)

type TableEntry struct {
	key   string
	value interface{}
	next  *TableEntry
}

type HashTable struct {
	table      []*TableEntry
	size       int
	count      int
	lowerBound float64
	upperBound float64
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		table:      make([]*TableEntry, size),
		size:       size,
		count:      0,
		lowerBound: 0.25, // resize when count <= 0.25 * size
		upperBound: 1,    // resize when count >= size
	}
}

func (ht *HashTable) simpleHash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ht.size
}

func (ht *HashTable) Insert(key string, value interface{}) {
	index := ht.simpleHash(key)
	entry := ht.table[index]
	for entry != nil {
		if entry.key == key {
			entry.value = value
			return
		}
		entry = entry.next
	}
	ht.table[index] = &TableEntry{key: key, value: value, next: ht.table[index]}
	ht.count++

	if float64(ht.count)/float64(ht.size) >= ht.upperBound {
		ht.Resize(2 * ht.size)
	}
}

func (ht *HashTable) Get(key string) interface{} {
	index := ht.simpleHash(key)
	entry := ht.table[index]
	for entry != nil {
		if entry.key == key {
			return entry.value
		}
		entry = entry.next
	}
	return nil
}

func (ht *HashTable) Remove(key string) interface{} {
	index := ht.simpleHash(key)
	entry := ht.table[index]
	var previous *TableEntry
	for entry != nil {
		if entry.key == key {
			if previous != nil {
				previous.next = entry.next
			} else {
				ht.table[index] = entry.next
			}
			ht.count--

			if float64(ht.count)/float64(ht.size) <= ht.lowerBound {
				ht.Resize(ht.size / 2)
			}

			return entry.value
		}
		previous, entry = entry, entry.next
	}
	return nil
}

func (ht *HashTable) Resize(newSize int) {
	newTable := NewHashTable(newSize)
	for _, entry := range ht.table {
		for entry != nil {
			newTable.Insert(entry.key, entry.value)
			entry = entry.next
		}
	}
	ht.table = newTable.table
	ht.size = newTable.size
}

func (ht *HashTable) Print() string {
	var result string
	for i := 0; i < ht.size; i++ {
		bucket := ""
		current := ht.table[i]
		for current != nil {
			bucket += " -> (" + current.key + ", " + fmt.Sprint(current.value) + ")"
			current = current.next
		}
		if bucket == "" {
			result += "NULL\n"
		} else {
			result += bucket[4:] + " -> NULL\n"
		}
	}
	return result
}
