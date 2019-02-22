package data

type FileType int

const (
	Data FileType = iota
	Sequences
	Test
	VariationSequences
	ZWJSequences
)

var AllDataTypes = []FileType{Data, Sequences, Test, VariationSequences, ZWJSequences}
