package data

import (
	"fmt"
	"regexp"
)

const (
	cp            = "[[:xdigit:]]{4,5}"
	singlePattern = "^(" + cp + ")"
	rangePattern  = "^(" + cp + ")[.]{2}(" + cp + ")"
	seqPattern    = "^" + cp + "(?: " + cp + ")+"
)

var RangeRegexp = getRegexp(rangePattern)
var SingleRegexp = getRegexp(singlePattern)
var SeqRegexp = regexp.MustCompile(seqPattern)

var legacySeqTypeRegexps = map[SeqType]*regexp.Regexp {
	Emoji_Flag_Sequence: regexp.MustCompile("flag for"),
	Emoji_Combining_Sequence: regexp.MustCompile("keycap"),
	Emoji_Modifier_Sequence: regexp.MustCompile("type-[0-9]"),
}

var regexpCache = map[string]*regexp.Regexp{}

func getRegexp(regexpStr string) *regexp.Regexp {
	if re, ok := regexpCache[regexpStr]; ok {
		return re
	}
	re := regexp.MustCompile(regexpStr)
	regexpCache[regexpStr] = re
	return re
}

func hasPropertyRegexp(property Property) *regexp.Regexp {
	regexpStr := fmt.Sprintf(";\\s+%v\\s*[;#]", property)
	return getRegexp(regexpStr)
}

func seqTypeRegexp(seqType SeqType) *regexp.Regexp {
	regexpStr := fmt.Sprintf(";\\s+%v\\s*[;#]", seqType)
	return getRegexp(regexpStr)
}

func legacySeqTypeRegexp(seqType SeqType) (*regexp.Regexp, bool) {
	re, ok := legacySeqTypeRegexps[seqType]
	return re, ok
}
