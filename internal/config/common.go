package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

const (
	DatabaseTypeMySQL  DatabaseType = "mysql"
	DatabaseTypeSQLite DatabaseType = "sqlite"
)

type Config interface {
	CoreServerConfig | NewsLetterGeneratorConfig
}

type DatabaseType string

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
		return ".temp_db?_foreign_keys=on"
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

func ParseConfig[T Config](configPath string) (T, error) {
	var config T

	fd, err := os.Open(configPath)
	defer func() {
		err := fd.Close()
		if err != nil {
			fmt.Println("failed to close config file", err)
		}
	}()

	if err != nil {
		return config, fmt.Errorf("failed to open config file: %w", err)
	}

	config, err = parseReader[T](fd)
	if err != nil {
		return config, err
	}

	return config, nil
}

func parseReader[T Config](r io.Reader) (T, error) {
	var c T

	dec := yaml.NewDecoder(r)
	if err := dec.Decode(&c); err != nil {
		return c, fmt.Errorf("failed to parse config: %w", err)
	}

	return c, nil
}
