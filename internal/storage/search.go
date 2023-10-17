package storage

import (
	"sort"
	"strings"
)

func SearchBytes(predicate func(byte Byte) bool) ([]Byte, error) {
	bytes, err := GetAllBytes()
	if err != nil {
		return []Byte{}, err
	}

	var filtered []Byte

	if predicate == nil {
		filtered = bytes
	} else {
		for _, byte := range bytes {
			if predicate(byte) {
				filtered = append(filtered, byte)
			}
		}
	}

	// Sort the files by name
	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].Title, filtered[j].Title) > 0
	})

	return filtered, err
}
