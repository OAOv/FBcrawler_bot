package task

func GetTarget(s string) string {
	switch s {
	case "palladium":
		return "PalladiumBoots.Taiwan"
	case "azurlane":
		return "azurlaneTW"
	}
	return "NotFound"
}
