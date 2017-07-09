package config

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config is a collection of configuration
type Config struct {
	Port string `envconfig:"port" default:"7007"`
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
