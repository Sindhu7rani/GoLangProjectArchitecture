package config

/*
environment settings
Load application configuration (port, DB details, etc.)

Config is loaded first because other layers depend on it
-->we can change values without touching the code
*/
type Config struct {
	Port string
}

func LoadConfig() Config {
	return Config{
		Port: "8080",
	}
}
