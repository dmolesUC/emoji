package test

import (
	"github.com/dmolesUC3/emoji/emoji/data"
	. "gopkg.in/check.v1"
	"unicode"
)

type DataSuite struct {

}

var _ = Suite(&DataSuite{})

func (s *DataSuite) TestEmojiRangeTable(c *C) {
	rt := data.Latest.EmojiRangeTable()
	c.Assert(rt, NotNil)

	c.Assert(rt.LatinOffset, Equals, 5)

	samples := []string {
		"㊙",		// 3299
		"〽",		// 303D
		"⬅⬆⬇",	// 2B05, 2B06, 2B07
		"🈯",		// 1F22F
		"🚀🚢🛅",	// 1F680, 1F6A2, 1F6C5
		"🛳",		// 1F6F3
		"🧀",		// 1F9C0
		"🦅🦋🦑",	// 1F985, 1F98B, 1F991
		"🧐🧛🧦",	// 1F9D0, 1F9DB, 1F9E6
		"🧧🧳🧿",	// 1F9E7, 1F9F3, 1F9FF
		"🪐🪒🪕",	// 1FA90, 1FA92, 1FA95
	}

	for _, sample := range samples {
		for _, r := range sample {
			inRange := unicode.In(r, rt)
			c.Assert(inRange, Equals, true, Commentf("expected %v (%X) to be in range, but was not", string(r), r))
		}
	}
}