package data

// SeqType represents a Unicode emoji sequence type. Note that prior to version 3.0,
// type information is not included in the sequence data files.
type SeqType string

const (
	Emoji_Combining_Sequence SeqType = "Emoji_Combining_Sequence"
	Emoji_Flag_Sequence SeqType = "Emoji_Flag_Sequence"
	Emoji_Modifier_Sequence SeqType = "Emoji_Modifier_Sequence"
	Emoji_Tag_Sequence SeqType = "Emoji_Tag_Sequence"
	Emoji_ZWJ_Sequence SeqType = "Emoji_ZWJ_Sequence"
)

// AllSeqTypes lists all Unicode emoji sequence types.
var AllSeqTypes = []SeqType{
	Emoji_Combining_Sequence,
	Emoji_Flag_Sequence,
	Emoji_Modifier_Sequence,
	Emoji_Tag_Sequence,
	Emoji_ZWJ_Sequence,
}

// String returns the type name as a string.
func (p SeqType) String() string {
	return string(p)
}
