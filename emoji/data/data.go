//go:generate go get -u -v github.com/go-bindata/go-bindata/...
//go:generate go-bindata -prefix ../internal/unicode.org/Public/emoji/ -ignore \.*html -ignore ReadMe\.txt -pkg data -o emoji_bindata.go ../internal/unicode.org/Public/emoji/...
// NOTE: installing go-bindata currently requires using `GO111MODULE=off go generate`

package data

var sourceByVersionAndType = map[Version]map[SourceType][]byte {
	V1: {
		Data: __10EmojiDataTxt,
	},
	V2: {
		Data: __20EmojiDataTxt,
		Sequences: __20EmojiSequencesTxt,
		ZWJSequences: __20EmojiZwjSequencesTxt,
	},
	V3: {
		Data: __30EmojiDataTxt,
		Sequences: __30EmojiSequencesTxt,
		ZWJSequences: __30EmojiSequencesTxt,
	},
	V4: {
		Data: __40EmojiDataTxt,
		Sequences: __40EmojiSequencesTxt,
		Test: __40EmojiTestTxt,
		ZWJSequences: __40EmojiSequencesTxt,
	},
	V5: {
		Data: __50EmojiDataTxt,
		Sequences: __50EmojiSequencesTxt,
		Test: __50EmojiTestTxt,
		VariationSequences: __50EmojiVariationSequencesTxt,
		ZWJSequences: __50EmojiVariationSequencesTxt,
	},
	V11: {
		Data: __110EmojiDataTxt,
		Sequences: __110EmojiSequencesTxt,
		Test: __110EmojiTestTxt,
		VariationSequences: __110EmojiVariationSequencesTxt,
		ZWJSequences: __110EmojiVariationSequencesTxt,
	},
	V12: {
		Data: __120EmojiDataTxt,
		Sequences: __120EmojiSequencesTxt,
		Test: __120EmojiTestTxt,
		VariationSequences: __120EmojiVariationSequencesTxt,
		ZWJSequences: __120EmojiVariationSequencesTxt,
	},
}