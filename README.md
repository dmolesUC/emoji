# emoji

A basic wrapper around the [Unicode.org emoji data files](http://unicode.org/Public/emoji/).

## Documentation

- [Variables](#variables)
- [func ParseRangeTable](#func-parserangetable)
- [type Property](#type-property)
   - [func (Property) String](#func-property-string)
- [type Version](#type-version)
   - [func (Version) FileBytes](#func-version-filebytes)
   - [func (Version) HasFile](#func-version-hasfile)
   - [func (Version) RangeTable](#func-version-rangetable)

### Variables

AllProperties lists all emoji Unicode properties 

```
var AllProperties = []Property{
    Emoji,
    Emoji_Presentation,
    Emoji_Modifier_Base,
    Emoji_Component,
    Extended_Pictographic,
}
```

AllVersions lists all emoji versions in
order

```
var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}
```

### func ParseRangeTable

```
func ParseRangeTable(property Property, data []byte) *unicode.RangeTable
```

ParseRangeTable parses the specified Unicode.org data file for
characters with the specified property, and returns a range table of
those characters.

Note that the range table reflects the ranges as defined in the source
files; ranges are guaranteed not to overlap, as per the RangeTable docs,
but adjacent ranges are not coalesced.

### type Property

Property represents a Unicode emoji property

```
type Property string

const (
    Emoji                 Property = "Emoji"
    Emoji_Presentation    Property = "Emoji_Presentation"
    Emoji_Modifier_Base   Property = "Emoji_Modifier_Base"
    Emoji_Component       Property = "Emoji_Component"
    Extended_Pictographic Property = "Extended_Pictographic"
)
```

#### func (Property) String

```
func (p Property) String() string
```

### type Version

Version represents an Emoji major release, e.g. V5 for Emoji version
5.0. Note that starting at Emoji version 11.0, the Emoji version is
synchronized to the corresponding Unicode version, so there are no
versions 6-10.

```
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

#### func (Version) FileBytes

```
func (v Version) FileBytes(t FileType) []byte
```

FileBytes returns the byte data of the Unicode.org source file of the
specified type for this version, e.g. V12.FileBytes(Sequences) returns
the contents of the file
<http://unicode.org/Public/emoji/12.0/emoji-sequences.txt>

#### func (Version) HasFile

```
func (v Version) HasFile(t FileType) bool
```

HasFile returns true if this version has a file of the specified type,
false otherwise. E.g., ZWJ (zero width joiner) sequences were introduced
only in Emoji version 2.0, test files in version 4.0, and variation
sequences in version
5.0.

#### func (Version) RangeTable

```
func (v Version) RangeTable(property Property) *unicode.RangeTable
```

RangeTable returns the Unicode range table for characters with the
specified property in this Emoji version. Note that the range table
reflects the ranges as defined in the source files from Unicode.org;
ranges are guaranteed not to overlap, as per the RangeTable docs, but
adjacent ranges are not coalesced.
