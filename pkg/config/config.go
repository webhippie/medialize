package config

// Photos defines the photos stuff.
type Photos struct {
	Source string `mapstructure:"source"`
	Target string `mapstructure:"target"`
	Rename bool   `mapstructure:"rename"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Config is a combination of all available configurations.
type Config struct {
	Photos Photos `mapstructure:"photos"`
	Logs   Logs   `mapstructure:"log"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
