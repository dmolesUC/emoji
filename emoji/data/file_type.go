package data

type SourceType int

const (
	Data SourceType = iota
	Sequences
	Test
	VariationSequences
	ZWJSequences
)

func (t SourceType) HasData(v Version) bool {
	return v.HasData(t)
}

func (t SourceType) Source(v Version) []byte {
	return v.Source(t)
}