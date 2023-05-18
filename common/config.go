package common

import (
	"github.com/pelletier/go-toml"
)

type Config struct {
	LogLevel string
	HttpPort int
}

const (
	DefaultLogLevel = "info"
	DefaultPort     = 8080
)

func NewConfig() Config {
	config, err := toml.LoadFile("config.toml")

	if err != nil {
		return Config{
			LogLevel: DefaultLogLevel,
			HttpPort: DefaultPort,
		}
	}
	return Config{
		LogLevel: getOrDefaultString(config, "log.level", DefaultLogLevel),
		HttpPort: getOrDefaultInt(config, "server.http_port", DefaultPort),
	}
}

func getOrDefaultString(config *toml.Tree, value, def string) string {
	switch config.Get(value).(type) {
	case string:
		return config.Get(value).(string)
	}

	return def
}

func getOrDefaultInt(config *toml.Tree, value string, def int) int {
	switch config.Get(value).(type) {
	case int:
		return config.Get(value).(int)
	case int64:
		return int(config.Get(value).(int64))
	}

	return def
}
