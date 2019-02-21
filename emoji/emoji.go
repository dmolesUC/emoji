//go:generate go get -u github.com/go-bindata/go-bindata
//go:generate go install github.com/go-bindata/go-bindata
//go:generate go-bindata -prefix internal/unicode.org/Public/emoji/ -ignore \.*html -ignore ReadMe\.txt -pkg emoji -o internal/emoji_bindata.go internal/unicode.org/Public/emoji/...

package emoji

