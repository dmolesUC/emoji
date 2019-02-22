//go:generate go get -u -v github.com/go-bindata/go-bindata/...
//go:generate go-bindata -prefix unicode.org/Public/emoji/ -ignore \.*html -ignore ReadMe\.txt -pkg data -o bindata.go unicode.org/Public/emoji/...
// NOTE: installing go-bindata currently requires using `GO111MODULE=off go generate`

package data

var getBytesByVersionAndType = map[int]map[FileType]func() ([]byte, error){
	1: {
		Data: _10EmojiDataTxtBytes,
	},
	2: {
		Data:         _20EmojiDataTxtBytes,
		Sequences:    _20EmojiSequencesTxtBytes,
		ZWJSequences: _20EmojiZwjSequencesTxtBytes,
	},
	3: {
		Data:         _30EmojiDataTxtBytes,
		Sequences:    _30EmojiSequencesTxtBytes,
		ZWJSequences: _30EmojiSequencesTxtBytes,
	},
	4: {
		Data:         _40EmojiDataTxtBytes,
		Sequences:    _40EmojiSequencesTxtBytes,
		Test:         _40EmojiTestTxtBytes,
		ZWJSequences: _40EmojiSequencesTxtBytes,
	},
	5: {
		Data:              _50EmojiDataTxtBytes,
		Sequences:         _50EmojiSequencesTxtBytes,
		Test:              _50EmojiTestTxtBytes,
		VariationSequences: _50EmojiVariationSequencesTxtBytes,
		ZWJSequences:      _50EmojiVariationSequencesTxtBytes,
	},
	11: {
		Data:              _110EmojiDataTxtBytes,
		Sequences:         _110EmojiSequencesTxtBytes,
		Test:              _110EmojiTestTxtBytes,
		VariationSequences: _110EmojiVariationSequencesTxtBytes,
		ZWJSequences:      _110EmojiVariationSequencesTxtBytes,
	},
	12: {
		Data:              _120EmojiDataTxtBytes,
		Sequences:         _120EmojiSequencesTxtBytes,
		Test:              _120EmojiTestTxtBytes,
		VariationSequences: _120EmojiVariationSequencesTxtBytes,
		ZWJSequences:      _120EmojiVariationSequencesTxtBytes,
	},
}
