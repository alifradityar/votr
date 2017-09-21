package config

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config is a collection of configuration
type Config struct {
	MysqlHost            string `envconfig:"mysql_host" default:"192.168.99.100"`
	MysqlUsername        string `envconfig:"mysql_username" default:"root"`
	MysqlPassword        string `envconfig:"mysql_password" default:"root-is-not-used"`
	MysqlDatabase        string `envconfig:"mysql_database" default:"votr"`
	MysqlConnectionLimit int    `envconfig:"mysql_connection_limit" default:"50"`
	RedisHost            string `envconfig:"redis_host" default:"192.168.99.100"`
	RedisPort            int    `envconfig:"redis_port" default:"6379"`
	Port                 string `envconfig:"port" default:"7007"`
}

var conf Config
var once sync.Once

// Get returns a config singleton
func Get() Config {
	once.Do(func() {
		err := envconfig.Process("", &conf)
		if err != nil {
			fmt.Printf("Can't load config: %v", err)
		}
	})

	return conf
}
