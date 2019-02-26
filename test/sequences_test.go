package test

import (
	. "github.com/dmolesUC3/emoji"
	. "github.com/dmolesUC3/emoji/data"
	. "gopkg.in/check.v1"
)

type SequencesSuite struct {
}

var _ = Suite(&SequencesSuite{})

// TODO: figure out what (if anything) was *added* in each version
var samplesByVersionAndType = map[Version]map[SeqType]string{
	V1: {
		Emoji_Combining_Sequence: "9⃣", // note no variation selector
		Emoji_Flag_Sequence:      "🇿🇼",
	},
	V2: {
		Emoji_Flag_Sequence:      "🇿🇼",
		Emoji_Combining_Sequence: "9⃣", // note no variation selector
		Emoji_Modifier_Sequence:  "🤘🏿",
	},
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

func (s *SequencesSuite) TestLegacySequences(c *C) {
	ok := true
	for _, t := range AllSeqTypes {
		for _, v := range []Version{V1, V2} {
			for _, s := range s.combinedSamples(t, v) {
				ix := index(v.Sequences(t), s)
				ok = ok && c.Check(ix, Not(Equals), -1, Commentf("expected %v sequences for %v to include %#v (%X), but did not", t, v, s, []rune(s)))
			}
		}
	}
	c.Assert(ok, Equals, true)
}

func (s *SequencesSuite) TestSequences(c *C) {
	ok := true
	for _, t := range AllSeqTypes {
		for _, v := range AllVersions {
			for _, s := range s.combinedSamples(t, v) {
				ix := index(v.Sequences(t), s)
				ok = ok && c.Check(ix, Not(Equals), -1, Commentf("expected %v sequences for %v to include %#v (%X), but did not", t, v, s, []rune(s)))
			}
		}
	}
	c.Assert(ok, Equals, true)
}

func (s *SequencesSuite) TestDisplayWidth(c *C) {
	ok := true
	for _, v := range AllVersions {
		for _, t := range AllSeqTypes {
			seqs := v.Sequences(t)
			for _, s := range seqs {
				w := DisplayWidth(s)
				ok = ok && c.Check(w, Equals, 1, Commentf("expected \"%v\" (%#v, %X) in %v (%v) to have length 1, but was %d", s, s, []rune(s), v, t, w))
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
