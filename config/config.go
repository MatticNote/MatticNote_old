package config

import (
	_ "embed"
	"errors"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

const FileName string = "matticnote_config.toml"

//go:embed matticnote_config.default.toml
var defaultConfig []byte

var Config *MatticNoteConfig

type (
	MNCDb struct {
		Address    string
		Port       uint
		User       string
		Password   string
		Name       string
		Sslmode    string
		MaxConnect int
	}

	MNCSrv struct {
		Address     string
		Port        uint
		Endpoint    []string
		CsrfSecret  string `toml:"csrf_secret"`
		CsrfSecure  bool   `toml:"csrf_secure"`
		OauthSecret string `toml:"oauth_secret"`
	}

	MNCMeta struct {
		InstanceName      string `toml:"instance_name"`
		MaintainerName    string `toml:"maintainer_name"`
		MaintainerContact string `toml:"maintainer_contact"`
		RepositoryUrl     string `toml:"repository_url"`
	}

	MNCRegister struct {
		Mode            string
		RecaptchaSite   string `toml:"recaptcha_site"`
		RecaptchaSecret string `toml:"recaptcha_secret"`
	}

	MatticNoteConfig struct {
		Database MNCDb
		Server   MNCSrv
		Meta     MNCMeta
		Register MNCRegister
	}
)

func LoadConfiguration() (*MatticNoteConfig, error) {
	cfgRaw, err := ioutil.ReadFile(FileName)
	if os.IsNotExist(err) {
		return nil, errors.New("configuration file was not exists. Please create them")
	}
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
	if len(cfg.Server.Endpoint) == 0 {
		return errors.New("validation error: server endpoint must not be empty")
	}
	if cfg.Server.CsrfSecret == "" {
		return errors.New("validation error: csrf secret must not be empty")
	}
	if cfg.Meta.InstanceName == "" {
		return errors.New("validation error: instance name must not be empty")
	}
	if cfg.Meta.MaintainerName == "" {
		return errors.New("validation error: maintainer name must not be empty")
	}
	if cfg.Meta.RepositoryUrl == "" {
		return errors.New("validation error: repository url must not be empty")
	}
	switch cfg.Register.Mode {
	case "open", "close":
	case "":
		return errors.New("validation error: register mode must not be empty")
	default:
		return errors.New("validation error: not valid register mode value")
	}
	return nil
}

func CreateDefaultConfiguration(override bool) error {
	_, err := os.Stat(FileName)
	if !override && !os.IsNotExist(err) {
		return errors.New("configuration file exists")
	}

	err = ioutil.WriteFile(FileName, defaultConfig, 0755)
	if err != nil {
		return err
	}

	return nil
}
