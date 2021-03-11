package config

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mgutz/ansi"
	"os"
	"path"
)

func NewDefaultSecrets() *Secrets {
	return &Secrets{
		Secrets: make(map[string]string),
	}
}

func (cfg *Secrets) Save() (err error) {
	// write
	prerr := func(err error) {
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Error reading secrets (", err.Error(), ")", ansi.Reset)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		prerr(err)
		return
	}

	cfgPath := path.Join(home, ".gme.secrets.toml")

	// create default config
	f, err := os.Create(cfgPath)
	if err != nil {
		prerr(err)
		return
	}

	var b bytes.Buffer
	enc := toml.NewEncoder(&b)
	if err = enc.Encode(cfg); err != nil {
		prerr(err)
		return
	}

	if _, err = f.Write(b.Bytes()); err != nil {
		prerr(err)
	}

	return
}

func (cfg *Secrets) Add(id, secret string) (err error) {
	if cfg.Secrets == nil {
		cfg.Secrets = make(map[string]string)
	}
	cfg.Secrets[id] = secret
	return cfg.Save()
}

func (cfg *Secrets) Del(id string) (err error) {
	if cfg.Secrets == nil {
		cfg.Secrets = make(map[string]string)
	}
	delete(cfg.Secrets, id)
	return cfg.Save()
}

func ReadSecrets() (cfg *Secrets) {
	// default
	cfg = NewDefaultSecrets()

	prerr := func(err error) {
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Error reading secrets (", err.Error(), ")", ansi.Reset)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		prerr(err)
		return
	}

	cfgPath := path.Join(home, ".gme.secrets.toml")
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		// create default config
		f, err := os.Create(cfgPath)
		if err != nil {
			prerr(err)
			return
		}

		var b bytes.Buffer
		enc := toml.NewEncoder(&b)
		if err = enc.Encode(cfg); err != nil {
			prerr(err)
			return
		}

		if _, err = f.Write(b.Bytes()); err != nil {
			prerr(err)
			return
		}
	}

	// read file
	if _, err = toml.DecodeFile(cfgPath, cfg); err != nil {
		prerr(err)
	}

	return
}
