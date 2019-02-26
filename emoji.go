package emoji

import (
	"fmt"
	. "github.com/dmolesUC3/emoji/data"
	"unicode"
)

// ------------------------------------------------------------
// Misc functions

func IsRegionalIndicator(r rune) bool {
	return r >= 0x1f1e6 && r <= 0x1f1ff
}

func IsEmojiSkinToneModifier(r rune) bool {
	return r >= 0x1f3fb && r <= 0x1f3ff
}

// DisplayWidth attempts to guess at the display width of a string containing
// emoji, taking into account variation selectors (0xFE00-0xFE0F), zero-width
// joins (0x200D), combining diacritical marks (0x20d0-0x20ff), flags,
// and skin tone modifiers.
func DisplayWidth(str string) int {
	width := 0
	runes := []rune(str)
	for i, r := range runes {
		if unicode.Is(unicode.Variation_Selector, r) {
			continue
		}
		if r >= 0x20d0 && r <= 0x20ff {
			// Combining Diacritical Marks for Symbols
			continue
		}
		// TODO: generally something smarter to identify sequences
		if i > 0 && IsRegionalIndicator(r) && IsRegionalIndicator(runes[i - 1]) {
			// TODO: make sure we don't collapse multiple flags
			continue
		}
		if IsEmojiSkinToneModifier(r) {
			// TODO: make sure we only do so when part of a sequence
			continue
		}
		if r == '\u200d' && len(runes) > i+1 {
			// ZWJ effectively "suppresses" the next character
			width -= 1
		} else {
			width += 1
		}
	}
	return width
}

// ------------------------------------------------------------
// Version type

// Version represents an Emoji major release, e.g. V5 for Emoji version 5.0.
// Note that starting at Emoji version 11.0, the Emoji version is synchronized
// to the corresponding Unicode version, so there are no versions 6-10.
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

// AllVersions lists all emoji versions in order.
var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}

// String returns this version as a string, e.g. V4.String() -> "Emoji 4.0"
func (v Version) String() string {
	return fmt.Sprintf("Emoji %d.0", int(v))
}

// HasFile returns true if this version has a file of the specified type, false
// otherwise. E.g., ZWJ (zero width joiner) sequences were introduced only in
// Emoji version 2.0, test files in version 4.0, and variation sequences in version
// 5.0.
func (v Version) HasFile(t FileType) bool {
	return t.HasData(int(v))
}

// FileBytes returns the byte data of the Unicode.org source file of the specified type
// for this version, e.g. V12.FileBytes(Sequences) returns the contents of the file
// http://unicode.org/Public/emoji/12.0/emoji-sequences.txt
func (v Version) FileBytes(t FileType) []byte {
	bytes, err := t.GetBytes(int(v))
	if err == nil {
		return bytes
	}
	return nil
}

// RangeTable returns the Unicode range table for characters with the specified property
// in this Emoji version. Note that the range table reflects the ranges as defined in the
// source files from Unicode.org; ranges are guaranteed not to overlap, as per the RangeTable
// docs, but adjacent ranges are not coalesced.
func (v Version) RangeTable(property Property) *unicode.RangeTable {
	var exists bool
	var rtsByProperty map[Property]*unicode.RangeTable
	if rtsByProperty, exists = rangeTables[v]; !exists {
		rtsByProperty = map[Property]*unicode.RangeTable{}
		rangeTables[v] = rtsByProperty
	}
	var rt *unicode.RangeTable
	if rt, exists = rtsByProperty[property]; !exists {
		rt = ParseRangeTable(property, v.FileBytes(Data))
		rtsByProperty[property] = rt
	}
	return rt
}

// Sequences returns the Unicode emoji sequences of the specified type in this Emoji version.
func (v Version) Sequences(seqType SeqType) []string {
	var exists bool
	var seqsByType map[SeqType][]string
	if seqsByType, exists = sequences[v]; !exists {
		seqsByType = map[SeqType][]string{}
		sequences[v] = seqsByType
	}
	var seqs []string
	if seqs, exists = seqsByType[seqType]; !exists {
		var parseSeq ParseSeq
		if v == V1 || v == V2 {
			parseSeq = ParseSequencesLegacy
		} else {
			parseSeq = ParseSequences
		}

		var fileType FileType
		if v == V1 {
			fileType = Data
		} else if seqType == Emoji_ZWJ_Sequence {
			fileType = ZWJSequences
		} else {
			fileType = Sequences
		}

		seqs = parseSeq(seqType, v.FileBytes(fileType))
		seqsByType[seqType] = seqs
	}
	return seqs
}

// ------------------------------------------------------------
// Unexported symbols

var rangeTables = map[Version]map[Property]*unicode.RangeTable{}

var sequences = map[Version]map[SeqType][]string{}
