package config

import (
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
	// "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

var (
	once   sync.Once
	config *Config
)

func MustLoad() *Config {
	once.Do(func() {
		var configPath string
		flag.StringVar(&configPath, "config", "", "a string")
		flag.Parse()
		if configPath == "" {
			configPath = os.Getenv("CONFIG_PATH")
			if configPath == "" {
				log.Fatal("-config flag | CONFIG_PATH is not set")
			}
		}

		// Read config from file
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("read config file error: %s", err)
		}
		defer file.Close()

		config = &Config{}
		if err := yaml.NewDecoder(file).Decode(config); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}
	})

	return config
}

// INFO: another recommended
// func MustLoad() *Config {
// 	var configPath string
// 	flag.StringVar(&configPath, "config", "", "a string")
// 	flag.Parse()
// 	if configPath == "" {
// 		configPath = os.Getenv("CONFIG_PATH")
// 		if configPath == "" {
// 			log.Fatal("-config flag | CONFIG_PATH is not set")
// 		}
// 	}
//
// 	// check file is exist
// 	if _, err := os.Stat(configPath); os.IsNotExist(err) {
// 		log.Fatalf("config file does not exist: %s", configPath)
// 	}
//
// 	var config Config
// 	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
// 		log.Fatalf("cannot read config: %s", err)
// 	}
//
// 	return &config
// }
