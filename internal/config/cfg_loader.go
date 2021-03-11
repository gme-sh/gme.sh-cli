package config

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mgutz/ansi"
	"os"
	"path"
)

func NewDefaultConfig() *Config {
	return &Config{
		ApiUrl:      "https://gme.sh",
		SaveSecrets: true,
	}
}

func ReadConfig() (cfg *Config) {
	// default
	cfg = NewDefaultConfig()

	prerr := func(err error) {
		fmt.Println(ansi.Red+"ERROR:", ansi.White+"Error reading config (", err.Error(), ")", ansi.Reset)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		prerr(err)
		return
	}

	cfgPath := path.Join(home, ".gme.config.toml")
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
