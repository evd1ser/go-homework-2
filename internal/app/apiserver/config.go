package apiserver

//General config for rest api
type Config struct {
	//Port for start api
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

//Should return default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
