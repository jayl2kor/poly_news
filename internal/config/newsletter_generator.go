package config

type NewsLetterGeneratorConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

type SmtpConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	AuthEmail    string `yaml:"authEmail"`
	AuthPassword string `yaml:"authPassword"`
}
