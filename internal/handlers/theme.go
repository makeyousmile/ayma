package handlers

type ThemeOption struct {
	Key   string
	Label string
}

const defaultTheme = "terra"

func themeOptions() []ThemeOption {
	return []ThemeOption{
		{Key: "terra", Label: "Терракота"},
		{Key: "coast", Label: "Прибрежный"},
		{Key: "forest", Label: "Лесной"},
		{Key: "graphite", Label: "Графит"},
		{Key: "sunrise", Label: "Рассвет"},
		{Key: "sky", Label: "Небо"},
		{Key: "amber", Label: "Янтарь"},
		{Key: "aqua", Label: "Аква"},
		{Key: "aqua-deep", Label: "Аква (глубокая)"},
		{Key: "aqua-mist", Label: "Аква (дымка)"},
		{Key: "bulma", Label: "Bulma"},
		{Key: "pure", Label: "Pure"},
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
