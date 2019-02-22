package data

import (
	"fmt"
)

type FileType int

const (
	Data FileType = iota
	Sequences
	Test
	VariationSequences
	ZWJSequences
)

var AllFileTypes = []FileType{Data, Sequences, Test, VariationSequences, ZWJSequences}

func (t FileType) String() string {
	switch t {
	case Data:
		return "Data"
	case Sequences:
		return "Sequences"
	case Test:
		return "Test"
	case VariationSequences:
		return "VariationSequences"
	case ZWJSequences:
		return "ZWJSequences"
	default:
		return ""
	}
}

func (t FileType) HasData(v int) bool {
	if getBytesByType, ok := getBytesByVersionAndType[v]; ok {
		_, ok = getBytesByType[t]
		return ok
	}
	return false
}

func (t FileType) GetBytes(v int) ([]byte, error) {
	if getBytesByType, ok := getBytesByVersionAndType[v]; ok {
		if getBytes, ok := getBytesByType[t]; ok {
			return getBytes()
		}
	}
	return nil, fmt.Errorf("no data of type %v for v %d", t, v)
}
