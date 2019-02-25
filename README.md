# emoji

A basic wrapper around the [Unicode.org emoji data files](http://unicode.org/Public/emoji/).

## Documentation

- [emoji](#emoji-1)
   - [Variables](#variables)
      - [var AllVersions](#var-allversions)
   - [Types](#types)
      - [type Version](#type-version)
         - [func (v Version) FileBytes](#func-v-version-filebytes)
         - [func (v Version) HasFile](#func-v-version-hasfile)
         - [func (v Version) RangeTable](#func-v-version-rangetable)
         - [func (v Version) Sequences](#func-v-version-sequences)
         - [func (v Version) String](#func-v-version-string)
- [data](#data)
   - [Variables](#variables-1)
      - [var AllFileTypes](#var-allfiletypes)
   - [Functions](#functions)
      - [func ParseRangeTable](#func-parserangetable)
      - [func ParseSequences](#func-parsesequences)
      - [func ParseSequencesLegacy](#func-parsesequenceslegacy)
      - [func ParseSequencesMatching](#func-parsesequencesmatching)
   - [Types](#types-1)
      - [type FileType](#type-filetype)
         - [func (t FileType) GetBytes](#func-t-filetype-getbytes)
         - [func (t FileType) HasData](#func-t-filetype-hasdata)
         - [func (t FileType) String](#func-t-filetype-string)
      - [type Property](#type-property)
         - [func (p Property) String](#func-p-property-string)
      - [type SeqType](#type-seqtype)
         - [func (p SeqType) String](#func-p-seqtype-string)

### `emoji`

#### Variables

##### var AllVersions

```
var AllVersions = []Version{V1, V2, V3, V4, V5, V11, V12}
```

AllVersions lists all emoji versions in order.

#### Types

##### type Version

```
type Version int
```

Version represents an Emoji major release, e.g. V5 for Emoji version
5.0. Note that starting at Emoji version 11.0, the Emoji version is
synchronized to the corresponding Unicode version, so there are no
versions 6-10.

```
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

###### func (v Version) FileBytes

```
func (v Version) FileBytes(t FileType) []byte
```

FileBytes returns the byte data of the Unicode.org source file of the
specified type for this version, e.g. V12.FileBytes(Sequences) returns
the contents of the file
http://unicode.org/Public/emoji/12.0/emoji-sequences.txt

###### func (v Version) HasFile

```
func (v Version) HasFile(t FileType) bool
```

HasFile returns true if this version has a file of the specified type,
false otherwise. E.g., ZWJ (zero width joiner) sequences were introduced
only in Emoji version 2.0, test files in version 4.0, and variation
sequences in version 5.0.

###### func (v Version) RangeTable

```
func (v Version) RangeTable(property Property) *unicode.RangeTable
```

RangeTable returns the Unicode range table for characters with the
specified property in this Emoji version. Note that the range table
reflects the ranges as defined in the source files from Unicode.org;
ranges are guaranteed not to overlap, as per the RangeTable docs, but
adjacent ranges are not coalesced.

###### func (v Version) Sequences

```
func (v Version) Sequences(seqType SeqType) []string
```

Sequences returns the Unicode emoji sequences of the specified type in
this Emoji version.

###### func (v Version) String

```
func (v Version) String() string
```

String returns this version as a string, e.g. V4.String() -> "Emoji 4.0"

### `data`

#### Variables

##### var AllFileTypes

```
var AllFileTypes = []FileType{Data, Sequences, Test_, VariationSequences, ZWJSequences}
```

AllFileTypes lists all file types.

```
var AllProperties = []Property{
    Emoji,
    Emoji_Presentation,
    Emoji_Modifier_Base,
    Emoji_Component,
    Extended_Pictographic,
}
```

AllProperties lists all Unicode emoji properties.

```
var AllSeqTypes = []SeqType{
    Emoji_Combining_Sequence,
    Emoji_Flag_Sequence,
    Emoji_Modifier_Sequence,
    Emoji_Tag_Sequence,
    Emoji_ZWJ_Sequence,
}
```

AllSeqTypes lists all Unicode emoji sequence types.

#### Functions

##### func ParseRangeTable

```
func ParseRangeTable(property Property, data []byte) *unicode.RangeTable
```

ParseRangeTable parses the specified Unicode.org data file for
characters with the specified property, and returns a range table of
those characters.

Note that the range table reflects the ranges as defined in the source
files; ranges are guaranteed not to overlap, as per the RangeTable docs,
but adjacent ranges are not coalesced.

##### func ParseSequences

```
func ParseSequences(seqType SeqType, data []byte) []string
```

ParseSequences parses emoji sequences of the specified type from the
specified data. Note that prior to version 3.0, type information is not
included in the sequence data files, and ParseSequencesLegacy should be
used (with the appropriate file) instead.

##### func ParseSequencesLegacy

```
func ParseSequencesLegacy(seqType SeqType, data []byte) []string
```

ParseSequencesLegacy parses emoji sequences for Emoji versions 1.0 and
2.0. Note that in Emoji 1.0, all sequences are in the main data file
(filetype Data); for Emoji 2.0, all sequences are in the main sequences
file (filetype Sequences), with no subfiles for variation sequences, ZWJ
sequences, etc.

##### func ParseSequencesMatching

```
func ParseSequencesMatching(re *regexp.Regexp, data []byte) []string
```

ParseSequencesMatching parses emoji sequences from data lines matching
the specified regexp.

#### Types

##### type FileType

```
type FileType int
```

FileType represents the type of a Unicode.org data file. Note that the
"Test" type is declared as "Test_" to avoid name collisions.

```
const (
    Data FileType = iota
    Sequences
    Test_
    VariationSequences
    ZWJSequences
)
```

###### func (t FileType) GetBytes

```
func (t FileType) GetBytes(v int) ([]byte, error)
```

HasData returns data of this type for the specified Emoji major version
has data of this type, or nil and an error if no such data exists for
this version.

###### func (t FileType) HasData

```
func (t FileType) HasData(v int) bool
```

HasData returns true if the specified Emoji major version has data of
this type, false otherwise.

###### func (t FileType) String

```
func (t FileType) String() string
```

String returns the file type as a string.

##### type Property

```
type Property string
```

Property represents a Unicode emoji property.

```
const (
    Emoji                 Property = "Emoji"
    Emoji_Presentation    Property = "Emoji_Presentation"
    Emoji_Modifier_Base   Property = "Emoji_Modifier_Base"
    Emoji_Component       Property = "Emoji_Component"
    Extended_Pictographic Property = "Extended_Pictographic"
)
```

###### func (p Property) String

```
func (p Property) String() string
```

String returns the property name as a string.

##### type SeqType

```
type SeqType string
```

SeqType represents a Unicode emoji sequence type. Note that prior to
version 3.0, type information is not included in the sequence data
files.

```
const (
    Emoji_Combining_Sequence SeqType = "Emoji_Combining_Sequence"
    Emoji_Flag_Sequence      SeqType = "Emoji_Flag_Sequence"
    Emoji_Modifier_Sequence  SeqType = "Emoji_Modifier_Sequence"
    Emoji_Tag_Sequence       SeqType = "Emoji_Tag_Sequence"
    Emoji_ZWJ_Sequence       SeqType = "Emoji_ZWJ_Sequence"
)
```

###### func (p SeqType) String

```
func (p SeqType) String() string
```

String returns the type name as a string.

