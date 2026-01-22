package handlers

type ThemeOption struct {
	Key   string
	Label string
}

const defaultTheme = "ayma"

func themeOptions() []ThemeOption {
	return []ThemeOption{
		{Key: "ayma", Label: "Айма"},
	}
}

func isThemeAllowed(key string) bool {
	for _, option := range themeOptions() {
		if option.Key == key {
			return true
		}
	}
	return false
}
