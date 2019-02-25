package test

import (
	. "github.com/dmolesUC3/emoji"
	. "github.com/dmolesUC3/emoji/data"
	. "gopkg.in/check.v1"
)

type SequencesSuite struct {
}

var _ = Suite(&SequencesSuite{})

var samplesByVersionAndType = map[Version]map[SeqType]string{
	V3: {
		Emoji_Flag_Sequence:      "🇿🇼",
		Emoji_Combining_Sequence: "9️⃣",
		Emoji_Modifier_Sequence:  "🤾🏿",
		Emoji_ZWJ_Sequence:       "👩‍👩‍👧‍👧",
	},
	V4: {
		Emoji_Flag_Sequence:      "🇿🇼",
		Emoji_Combining_Sequence: "9️⃣",
		Emoji_Modifier_Sequence:  "🤾🏿",
		Emoji_ZWJ_Sequence:       "👁️‍🗨️",
	},
	V5: {
		Emoji_Flag_Sequence:     "🇿🇼",
		Emoji_Tag_Sequence:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		Emoji_Modifier_Sequence: "🧝🏿",
		Emoji_ZWJ_Sequence:      "👁️‍🗨️",
	},
	V11: {
		Emoji_Flag_Sequence:     "🇿🇼",
		Emoji_Tag_Sequence:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		Emoji_Modifier_Sequence: "🧝🏿",
		Emoji_ZWJ_Sequence:      "👁️‍🗨️",
	},
	V12: {
		Emoji_Modifier_Sequence: "🧝🏿",
		Emoji_Flag_Sequence:     "🇿🇼",
		Emoji_Tag_Sequence:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		Emoji_ZWJ_Sequence:      "👁️‍🗨️",
	},
}

func (s *SequencesSuite) combinedSamples(seqType SeqType, v Version) []string {
	var combined []string
	for _, v2 := range AllVersions {
		if v2 >= v {
			break
		}
		if samplesByType, ok := samplesByVersionAndType[v]; ok {
			if sample, ok := samplesByType[seqType]; ok {
				combined = append(combined, sample)
			}
		}
	}
	return combined
}

func (s *SequencesSuite) TestSequences(c *C) {
	ok := true
	types := AllSeqTypes
	versions := AllVersions

	types = []SeqType{Emoji_ZWJ_Sequence}
	versions = []Version{V5}

	for _, t := range types {
		for _, v := range versions {
			for _, s := range s.combinedSamples(t, v) {
				ix := index(v.Sequences(t), s)
				ok = ok && c.Check(ix, Not(Equals), -1, Commentf("expected %v sequences for %v to include %#v (%X), but did not", t, v, s, []rune(s)))
			}
		}
	}
	c.Assert(ok, Equals, true)
}

func index(strings []string, str string) int {
	for i, s := range strings {
		if s == str {
			return i
		}
	}
	return -1
}
