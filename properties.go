package emoji

// Property represents a Unicode emoji property
type Property string

//noinspection GoSnakeCaseUsage,GoNameStartsWithPackageName
const (
	Emoji                 Property = "Emoji"
	Emoji_Presentation    Property = "Emoji_Presentation"
	Emoji_Modifier_Base   Property = "Emoji_Modifier_Base"
	Emoji_Component       Property = "Emoji_Component"
	Extended_Pictographic Property = "Extended_Pictographic"
)

// AllProperties lists all emoji Unicode properties
//noinspection GoUnusedGlobalVariable
var AllProperties = []Property{
	Emoji,
	Emoji_Presentation,
	Emoji_Modifier_Base,
	Emoji_Component,
	Extended_Pictographic,
}

func (p Property) String() string {
	return string(p)
}
