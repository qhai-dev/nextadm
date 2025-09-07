package mysql

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Timeout      time.Duration `yaml:"timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	User         string        `yaml:"user"`
	Password     string        `yaml:"password"`
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	Name         string        `yaml:"name"`
}

func (cfg *Config) dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	args := []string{
		"charset=utf8mb4",
		"parseTime=True",
		"loc=Local",
		"timeout=" + cfg.Timeout.String(),
		"readTimeout=" + cfg.ReadTimeout.String(),
		"writeTimeout=" + cfg.WriteTimeout.String(),
	}

	return fmt.Sprintf("%s?%s", dsn, strings.Join(args, "&"))
}
