package config

import (
	"errors"
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

//goland:noinspection GoUnusedGlobalVariable
var Config *MatticNoteConfig

type MNCDbConfig struct {
	Address  string
	Port     uint
	User     string
	Password string
	Name     string
	Sslmode  string
}

type MNCSrvConfig struct {
	Address string
	Port    uint
}

type MatticNoteConfig struct {
	Database MNCDbConfig
	Server   MNCSrvConfig
}

func LoadConfiguration() (*MatticNoteConfig, error) {
	cfgRaw, err := ioutil.ReadFile("matticnote_config.toml")
	if err != nil {
		return nil, err
	}
	cfg := &MatticNoteConfig{}
	err = toml.Unmarshal(cfgRaw, cfg)
	if err != nil {
		return nil, err
	}
	if err := ValidateConfiguration(cfg); err != nil {
		return nil, err
	}
	Config = cfg
	return cfg, nil
}

func ValidateConfiguration(cfg *MatticNoteConfig) error {
	if cfg.Database.Port == 0 {
		return errors.New("validation error: database port must not be 0")
	}
	if cfg.Server.Port == 0 {
		return errors.New("validation error: server port must not be 0")
	}
	if cfg.Database.Address == "" {
		return errors.New("validation error: database address must not be empty")
	}
	if cfg.Server.Address == "" {
		return errors.New("validation error: server address must not be empty")
	}
	return nil
}

func CreateDefaultConfiguration() error {
	return nil
}
