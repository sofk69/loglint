// pkg/analyzer/config/config.go
package config

type Config struct {
	DisableLowercase bool     `mapstructure:"disable-lowercase"`
	DisableEnglish   bool     `mapstructure:"disable-english"`
	DisableSpecial   bool     `mapstructure:"disable-special-chars"`
	DisableSensitive bool     `mapstructure:"disable-sensitive"`
	ExtraSensitive   []string `mapstructure:"extra-sensitive-keywords"`
}
