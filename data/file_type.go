package data

import (
	"fmt"
)

// FileType represents the type of a Unicode.org data file. Note that the "Test"
// type is declared as "Test_" to avoid name collisions.
type FileType int

const (
	Data FileType = iota
	Sequences
	Test_
	VariationSequences
	ZWJSequences
)

// AllFileTypes lists all file types.
var AllFileTypes = []FileType{Data, Sequences, Test_, VariationSequences, ZWJSequences}

// String returns the file type as a string.
func (t FileType) String() string {
	switch t {
	case Data:
		return "Data"
	case Sequences:
		return "Sequences"
	case Test_:
		return "Test"
	case VariationSequences:
		return "VariationSequences"
	case ZWJSequences:
		return "ZWJSequences"
	default:
		return ""
	}
}

// HasData returns true if the specified Emoji major version has data of this type, false otherwise.
func (t FileType) HasData(v int) bool {
	if getBytesByType, ok := getBytesByVersionAndType[v]; ok {
		_, ok = getBytesByType[t]
		return ok
	}
	return false
}

// HasData returns data of this type for the specified Emoji major version has data of this type,
// or nil and an error if no such data exists for this version.
func (t FileType) GetBytes(v int) ([]byte, error) {
	if getBytesByType, ok := getBytesByVersionAndType[v]; ok {
		if getBytes, ok := getBytesByType[t]; ok {
			return getBytes()
		}
	}
	return nil, fmt.Errorf("no data of type %v for v %d", t, v)
}
