package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"runtime"
)

type JWT struct {
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Ttl    int    `yaml:"ttl"`
}

type Config struct {
	Port string `yaml:"port"`
	JWT  JWT    `yaml:"jwt"`
}

func New() (*Config, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("unable to get config file directory")
	}
	f, err := os.Open(path.Join(path.Dir(filename), "../config/config.yaml"))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	cfg := &Config{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
