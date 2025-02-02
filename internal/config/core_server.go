package config

type CoreServerConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Debug  bool   `yaml:"debug"`
	Listen string `yaml:"listen"`
}
