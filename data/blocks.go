package data

import (
	"unicode"
)

// These variables have type *unicode.RangeTable.
var (
	CombiningDiacritical  = _CombiningDiacritical	// Unicode block "Combining Diacritical Marks for Symbols"
	RegionalIndicator     = _RegionalIndicator		// Subset of "Enclosed Alphanumeric Supplement" used for flag regional indicators
	EmojiSkinToneModifier = _EmojiSkinToneModifier	// A.k.a. "EMOJI MODIFIER FITZPATRICK TYPE-(1-2|3|4|5|6)"
	Tag                   = _Tag					// Unicode block "Tags" used for subnational flag sequences
)

var _CombiningDiacritical = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x20d0, 0x20ff, 1},
	},
}

var _RegionalIndicator = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x1f1e6, 0x1f1ff, 1},
	},
}

var _EmojiSkinToneModifier = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0x1f3fb, 0x1f3ff, 1},
	},
}

var _Tag = &unicode.RangeTable{
	R32: []unicode.Range32{
		{0xe0000, 0xe007f, 1},
	},
}
