package emoji

import (
	. "github.com/dmolesUC3/emoji/pkg/data"
	. "github.com/dmolesUC3/emoji/pkg/properties"
	"unicode"
)

type Version int

const (
	V1 Version = 1
	V2 Version = 2
	V3 Version = 3
	V4 Version = 4
	V5 Version = 5
	// Starting at V11, Emoji version = Unicode version
	V11 Version = 11
	V12 Version = 12

	Latest = V12
)

var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}

var rangeTables = map[Version]map[Property]*unicode.RangeTable{}

func (v Version) HasData(t DataType) bool {
	return v.Source(t) != nil
}

func (v Version) Source(t DataType) []byte {
	if dataByVersion, ok := dataByVersionAndType[v]; ok {
		if d, ok := dataByVersion[t]; ok {
			return d
		}
	}
	return nil
}

func (v Version) RangeTable(property Property) *unicode.RangeTable {
	var exists bool
	var rtsByProperty map[Property]*unicode.RangeTable
	if rtsByProperty, exists = rangeTables[v]; !exists {
		rtsByProperty = map[Property]*unicode.RangeTable{}
		rangeTables[v] = rtsByProperty
	}
	var rt *unicode.RangeTable
	if rt, exists = rtsByProperty[property]; !exists {
		rt = ParseRangeTable(property, v.Source(Data))
		rtsByProperty[property] = rt
	}
	return rt
}
