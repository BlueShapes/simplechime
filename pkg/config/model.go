package config

type Config struct {
	manualTime []string
}

func GetDefaultConfig() Config {
	return Config{
		manualTime: []string{"9:00", "12:00", "18:45"},
	}
}
