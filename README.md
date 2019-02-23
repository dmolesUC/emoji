# emoji

A basic wrapper around the [Unicode.org emoji data files](http://unicode.org/Public/emoji/).

## Documentation

- [emoji](#emoji-1)
   - [Variables](#variables)
   - [type Version](#type-version)
      - [func (Version) FileBytes](#func-version-filebytes)
      - [func (Version) HasFile](#func-version-hasfile)
      - [func (Version) RangeTable](#func-version-rangetable)
- [emoji/data](#emojidata)
   - [Variables](#variables-1)
   - [func ParseRangeTable](#func-parserangetable)
   - [type FileType](#type-filetype)
      - [func (FileType) GetBytes](#func-filetype-getbytes)
      - [func (FileType) HasData](#func-filetype-hasdata)
      - [func (FileType) String](#func-filetype-string)
   - [type Property](#type-property)
      - [func (Property) String](#func-property-string)


### `emoji`

#### Variables

AllVersions lists all emoji versions in order.

```go
var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}
```

#### type Version

Version represents an Emoji major release, e.g. V5 for Emoji version
5.0. Note that starting at Emoji version 11.0, the Emoji version is
synchronized to the corresponding Unicode version, so there are no
versions 6-10.

```go
type Version int

const (
    V1 Version = 1
    V2 Version = 2
    V3 Version = 3
    V4 Version = 4
    V5 Version = 5
    // Starting at V11, Emoji version = Unicode version
    V11 Version = 11
    V12 Version = 12

    Latest = V12
)
```

##### func (Version) FileBytes

```go
func (v Version) FileBytes(t FileType) []byte
```

FileBytes returns the byte data of the Unicode.org source file of the
specified type for this version, e.g. V12.FileBytes(Sequences) returns
the contents of the file
<http://unicode.org/Public/emoji/12.0/emoji-sequences.txt>

##### func (Version) HasFile

```go
func (v Version) HasFile(t FileType) bool
```

HasFile returns true if this version has a file of the specified type,
false otherwise. E.g., ZWJ (zero width joiner) sequences were introduced
only in Emoji version 2.0, test files in version 4.0, and variation
sequences in version
5.0.

##### func (Version) RangeTable

```go
func (v Version) RangeTable(property Property) *unicode.RangeTable
```

RangeTable returns the Unicode range table for characters with the
specified property in this Emoji version. Note that the range table
reflects the ranges as defined in the source files from Unicode.org;
ranges are guaranteed not to overlap, as per the RangeTable docs, but
adjacent ranges are not coalesced.

### `emoji/data`

#### Variables

AllProperties lists all emoji Unicode properties 

```go
var AllProperties = []Property{
    Emoji,
    Emoji_Presentation,
    Emoji_Modifier_Base,
    Emoji_Component,
    Extended_Pictographic,
}
```

#### func ParseRangeTable

```go
func ParseRangeTable(property Property, data []byte) *unicode.RangeTable
```

ParseRangeTable parses the specified Unicode.org data file for
characters with the specified property, and returns a range table of
those characters.

Note that the range table reflects the ranges as defined in the source
files; ranges are guaranteed not to overlap, as per the RangeTable docs,
but adjacent ranges are not coalesced.

#### type FileType

```go
type FileType int
```

FileType represents the type of a Unicode.org data file.

```go
const (
	Data FileType = iota
	Sequences
	Test
	VariationSequences
	ZWJSequences
)
```

##### func (FileType) GetBytes

```go
func (t FileType) GetBytes(v int) ([]byte, error)
```
HasData returns data of this type for the specified Emoji major version has data
of this type, or nil and an error if no such data exists for this version.

##### func (FileType) HasData

```go
func (t FileType) HasData(v int) bool
```

HasData returns true if the specified Emoji major version has data of this type,
false otherwise.

##### func (FileType) String

```go
func (t FileType) String() string
```

String returns the file type as a string.

#### type Property

```go
type Property string
```

Property represents a Unicode emoji property.

```go
const (
	Emoji                 Property = "Emoji"
	Emoji_Presentation    Property = "Emoji_Presentation"
	Emoji_Modifier_Base   Property = "Emoji_Modifier_Base"
	Emoji_Component       Property = "Emoji_Component"
	Extended_Pictographic Property = "Extended_Pictographic"
)
```

##### func (Property) String

```go
func (p Property) String() string
```

String returns the property name as a string.
