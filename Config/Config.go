package Config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	config *Config
)

type Config struct {
	Db struct {
		Host     string `yaml:"host"`
		Port     int64  `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	Log LogConfig
}

func readConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}
	fmt.Println(config)

	return &conf, nil
}

func init() {
	dir, _ := os.Getwd()
	filename := fmt.Sprintf("%s%s", dir, "/conf/app.yaml")
	var err error = nil
	config, err = readConfig(filename)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}
