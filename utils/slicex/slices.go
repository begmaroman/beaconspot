package slicex

import "bytes"

func RemoveUint64Duplicate(intSlice []uint64) []uint64 {
	keys := make(map[uint64]bool)
	list := []uint64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ContainsString verifies if a bytes exists on a slice
func ContainsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsBytes verifies if a bytes exists on a slice
func ContainsBytes(slice [][]byte, value []byte) bool {
	for _, v := range slice {
		if bytes.Equal(v, value) {
			return true
		}
	}
	return false
}
