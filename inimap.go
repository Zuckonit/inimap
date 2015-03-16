package inimap

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type SubConfig map[string]string
type Config map[string]SubConfig

func (cfg *Config) Set(sec, key, value string) {
	(*cfg)[sec][key] = value
}

func (cfg *Config) Get(sec, key string) string {
	return (*cfg)[sec][key]
}

func (cfg *Config) Has(sec, key string) bool {
	_, ok := (*cfg)[sec][key]
	return ok
}

func (cfg *Config) GetInt(sec, key string) (int, error) {
	v := cfg.Get(sec, key)
	return strconv.Atoi(v)
}

func (cfg *Config) GetBool(sec, key string) (bool, error) {
	v := cfg.Get(sec, key)
	return strconv.ParseBool(v)
}

func (cfg *Config) GetSlice(sec, key, sep string) []string {
	return strings.Split(cfg.Get(sec, key), sep)
}

func isFileExisted(filename string) bool {
	fp, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !fp.IsDir()
}

func ReadIO(iniString []byte) (*Config, error) {
	cfg := make(Config)
	lines := strings.Split(string(iniString), "\n")
	var curSec string
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}
		if !strings.Contains(line, "=") && strings.HasPrefix(line, "[") &&
			strings.HasSuffix(line, "]") {
			//section
			sec := strings.TrimLeft(line, "[")
			sec = strings.TrimRight(sec, "]")
			sec = strings.TrimSpace(sec)
			if sec == "" {
				continue
			}
			if _, ok := cfg[sec]; !ok {
				cfg[sec] = make(map[string]string)
				curSec = sec
			}
		} else {
			s := strings.SplitN(line, "=", 2)
			if curSec == "" {
				continue
			}
			cfg[curSec][strings.TrimSpace(s[0])] = strings.TrimSpace(s[1])
		}
	}
	if len(cfg) == 0 {
		return nil, errors.New("ini is invalid")
	}
	return &cfg, nil
}

func ReadFile(cfgPath string) (*Config, error) {
	if !isFileExisted(cfgPath) {
		return nil, errors.New("config file not existed")
	}
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	return ReadIO(data)
}
