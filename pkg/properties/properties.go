package properties

type Property string

const (
	Emoji                 Property = "Emoji"
	Emoji_Presentation    Property = "Emoji_Presentation"
	Emoji_Modifier_Base   Property = "Emoji_Modifier_Base"
	Emoji_Component       Property = "Emoji_Component"
	Extended_Pictographic Property = "Extended_Pictographic"
)

func (p Property) String() string {
	return string(p)
}
