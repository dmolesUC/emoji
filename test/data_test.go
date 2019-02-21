package test

import (
	"github.com/dmolesUC3/emoji/emoji/data"
	. "gopkg.in/check.v1"
)

type DataSuite struct {

}

var _ = Suite(&DataSuite{})

func (s *DataSuite) TestRangeTable(c *C) {
	rt := data.Latest.RangeTable()
	c.Assert(rt, NotNil)

	// TODO: don't parse repeats
	//  - figure out what we want to parse:
	//    - only Emoji?
	c.Assert(rt.LatinOffset, Equals, 5)
}