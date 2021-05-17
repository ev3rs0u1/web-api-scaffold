package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"web-api-scaffold/internal/pkg/dirutil"
)

const (
	configFile      = "conf.yaml"
	serverLogFile   = "binfs-logs/server.log"
	databaseLogFile = "binfs-logs/database.log"
)

var (
	cfg = configurator{
		Server: &server{
			Host: "0.0.0.0",
			Port: 10985,
		},
		Database: &database{
			Driver:   "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Name:     "binfile",
			Username: "root",
			Password: "yt1024!@",
		},
		Logfile: &logfile{
			Mode:     "release",
			Server:   dirutil.JoinCurrentDir(serverLogFile),
			Database: dirutil.JoinCurrentDir(databaseLogFile),
		},
	}
)

func LoadFile() {
	if file, err := os.Open(configFile); err == nil {
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		decoder.KnownFields(true)

		if err := decoder.Decode(&cfg); err != nil {
			panic("解析配置文件时出错: " + err.Error())
		}
	}
}

func Instance() *configurator {
	return &cfg
}

func IsReleaseLogMode() bool {
	return cfg.Logfile.Mode == "release"
}
