package config

import (
	"errors"
	"fmt"
)

// configurator combine every child config
type configurator struct {
	Server   *server   `yaml:"server"`
	Database *database `yaml:"database"`
	Logfile  *logfile  `yaml:"logfile"`
}

// server define config format for http service
type server struct {
	Host string `yaml:"host"`
	Port uint32 `yaml:"port"`
}

// database represent database config
type database struct {
	// driver can be sqlite3, mysql or postgres
	Driver string `yaml:"driver"`

	// mysql and postgres database config
	// database server host
	Host string `yaml:"host"`

	// database username
	Username string `yaml:"username"`

	// database password
	Password string `yaml:"password"`

	// database name
	Name string `yaml:"name"`

	// database server port
	Port uint32 `yaml:"port"`
}

type logfile struct {
	Mode     string `yaml:"mode"`
	Server   string `yaml:"server"`
	Database string `yaml:"database"`
}

// DSN generate dsn based on db driver
func (d *database) DSN() (string, error) {
	switch d.Driver {
	case "mysql":
		format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2fShanghai"
		return fmt.Sprintf(format, d.Username, d.Password, d.Host, d.Port, d.Name), nil
	case "postgres":
		format := "host=%s port=%d user=%s dbname=%s password=%s"
		return fmt.Sprintf(format, d.Host, d.Port, d.Username, d.Name, d.Password), nil
	default:
		return "", errors.New("unsupported database driver")
	}
}
