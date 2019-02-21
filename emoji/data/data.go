//go:generate go get -u -v github.com/go-bindata/go-bindata/...
//go:generate go-bindata -prefix ../internal/unicode.org/Public/emoji/ -ignore \.*html -ignore ReadMe\.txt -pkg data -o emoji_bindata.go ../internal/unicode.org/Public/emoji/...
// NOTE: installing go-bindata currently requires using `GO111MODULE=off go generate`

package data

func mustRead(data []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return data
} 

var sourceByVersionAndType = map[Version]map[SourceType][]byte {
	V1: {
		Data: mustRead(_10EmojiDataTxtBytes()),
	},
	V2: {
		Data: mustRead(_20EmojiDataTxtBytes()),
		Sequences: mustRead(_20EmojiSequencesTxtBytes()),
		ZWJSequences: mustRead(_20EmojiZwjSequencesTxtBytes()),
	},
	V3: {
		Data: mustRead(_30EmojiDataTxtBytes()),
		Sequences: mustRead(_30EmojiSequencesTxtBytes()),
		ZWJSequences: mustRead(_30EmojiSequencesTxtBytes()),
	},
	V4: {
		Data: mustRead(_40EmojiDataTxtBytes()),
		Sequences: mustRead(_40EmojiSequencesTxtBytes()),
		Test: mustRead(_40EmojiTestTxtBytes()),
		ZWJSequences: mustRead(_40EmojiSequencesTxtBytes()),
	},
	V5: {
		Data: mustRead(_50EmojiDataTxtBytes()),
		Sequences: mustRead(_50EmojiSequencesTxtBytes()),
		Test: mustRead(_50EmojiTestTxtBytes()),
		VariationSequences: mustRead(_50EmojiVariationSequencesTxtBytes()),
		ZWJSequences: mustRead(_50EmojiVariationSequencesTxtBytes()),
	},
	V11: {
		Data: mustRead(_110EmojiDataTxtBytes()),
		Sequences: mustRead(_110EmojiSequencesTxtBytes()),
		Test: mustRead(_110EmojiTestTxtBytes()),
		VariationSequences: mustRead(_110EmojiVariationSequencesTxtBytes()),
		ZWJSequences: mustRead(_110EmojiVariationSequencesTxtBytes()),
	},
	V12: {
		Data: mustRead(_120EmojiDataTxtBytes()),
		Sequences: mustRead(_120EmojiSequencesTxtBytes()),
		Test: mustRead(_120EmojiTestTxtBytes()),
		VariationSequences: mustRead(_120EmojiVariationSequencesTxtBytes()),
		ZWJSequences: mustRead(_120EmojiVariationSequencesTxtBytes()),
	},
}