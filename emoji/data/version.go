package data

type Version int

const (
	V1 Version = 1
	V2 Version = 2
	V3 Version = 3
	V4 Version = 4
	V5 Version = 5
	// no V6-V10
	V11 Version = 11
	V12 Version = 12

	Latest = V12
)

func (v Version) HasData(t SourceType) bool {
	return v.Source(t) != nil
}

func (v Version) Source(t SourceType) []byte {
	if sourceByVersion, ok := sourceByVersionAndType[v]; ok {
		if data, ok := sourceByVersion[t]; ok {
			return data
		}
	}
	return nil
}