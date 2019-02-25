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
	for _, t := range AllSeqTypes {
		for _, v := range AllVersions {
			for _, s := range s.combinedSamples(t, v) {
				found := false
				for _, seq := range v.Sequences(t) {
					if s == seq {
						found = true
						break
					}
				}
				ok = ok && c.Check(found, Equals, true, Commentf("expected %v sequences for %v to include %#v, but did not", t, v, s))
			}
		}
	}
	c.Assert(ok, Equals, true)
}
