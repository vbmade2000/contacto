package config

import (
	"github.com/pelletier/go-toml"
	"io"
	"io/ioutil"
	"os"
)

type DBProvider struct {
	Provider string
	Port     int
	Host     string
}
type Config struct {
	ServerPort     int
	Host           string
	DBProvider     string
	DBProviderConf []DBProvider
}

func WriteConfig(config Config, out io.Writer) error {
	enc := toml.NewEncoder(out)
	if err := enc.Encode(config); err != nil {
		return err
	}
	_, err := out.Write([]byte{'\n'})
	return err
}

func GetDefaultConfig() Config {
	// Database configuration
	dbProviderConf := DBProvider{}
	dbProviderConf.Provider = "sqlite3"
	dbProviderConf.Host = "contacto.db"
	
	conf := Config{}
	conf.ServerPort = 8080
	conf.Host = "localhost"
	conf.DBProvider = "sqlite3"
	conf.DBProviderConf = []DBProvider{dbProviderConf}
	return conf
}
func ReadConfig(in io.Reader) (Config, error) {
	config := new(Config)
	rawData, err := ioutil.ReadAll(in)
	if err != nil {
		return *config, err
	}
	err = toml.Unmarshal(rawData, config)
	return *config, err
}

func ConfigExists() bool {
	return FileExists("contacto.conf")
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
