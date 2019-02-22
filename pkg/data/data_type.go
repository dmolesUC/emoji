package data

type DataType int

const (
	Data DataType = iota
	Sequences
	Test
	VariationSequences
	ZWJSequences
)

var AllDataTypes = []DataType{Data, Sequences, Test, VariationSequences, ZWJSequences}
