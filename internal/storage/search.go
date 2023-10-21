package storage

import (
	"sort"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type WeightedByte struct {
	Byte
	Weight int
}

func SearchBytes(predicate func(byte Byte) any) ([]Byte, error) {
	bytes, err := GetAllBytes()
	if err != nil {
		return []Byte{}, err
	}

	var filtered []Byte

	if predicate == nil {
		filtered = bytes
	} else {
		var weighted []WeightedByte
		var highestWeight int

		for _, byte := range bytes {
			result := predicate(byte)
			var weightedByte WeightedByte

			if result == true {
				weightedByte = WeightedByte{byte, 1}
			} else if weight, ok := result.(int); ok && weight > 0 {
				weightedByte = WeightedByte{byte, weight}
			}

			weighted = append(weighted, weightedByte)
			highestWeight = max(highestWeight, weightedByte.Weight)
		}

		// Pick the bytes with the highest weight
		for _, weightedByte := range weighted {
			if weightedByte.Weight == highestWeight {
				filtered = append(filtered, weightedByte.Byte)
			}
		}
	}

	// Sort the files by name
	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].Title, filtered[j].Title) > 0
	})

	return filtered, err
}
