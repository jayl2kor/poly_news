package config

import (
	"fmt"
	"io"
	"os"
	"poly_news/cmd/core-server/flag"

	"gopkg.in/yaml.v3"
)

const (
	DatabaseTypeMySQL  DatabaseType = "mysql"
	DatabaseTypeSQLite DatabaseType = "sqlite"
)

type DatabaseType string

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Debug  bool   `yaml:"debug"`
	Listen string `yaml:"listen"`
}

type DatabaseConfig struct {
	Type     DatabaseType `yaml:"type"`
	Username string       `yaml:"username"`
	Password string       `yaml:"password"`
	DBName   string       `yaml:"dbname"`
	Host     string       `yaml:"host"`
	Port     int          `yaml:"port"`

	SQLiteFilePath string `yaml:"sqliteFilePath"`
}

func (d DatabaseConfig) DSN() string {
	switch d.Type {
	case "sqlite":
		return "temp_db?_foreign_keys=on"
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			d.Username,
			d.Password,
			d.Host,
			d.Port,
			d.DBName,
		)
	}
	return ""
}

func Parse(flags flag.CoreServerFlag) (Config, error) {
	fd, err := os.Open(flags.ConfigPath)
	defer func() {
		err := fd.Close()
		if err != nil {
			fmt.Println("failed to close configfile", err)
		}
	}()

	if err != nil {
		return Config{}, fmt.Errorf("failed to open configfile: %w", err)
	}

	config, err := parseReader(fd)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func parseReader(r io.Reader) (Config, error) {
	c := Config{}

	dec := yaml.NewDecoder(r)
	if err := dec.Decode(&c); err != nil {
		return Config{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return c, nil
}
