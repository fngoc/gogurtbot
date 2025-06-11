package viper

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

var configName string
var configType string
var configPaths []string
var data map[string]interface{}

func SetConfigName(name string) { configName = name }
func SetConfigType(t string)    { configType = t }
func AddConfigPath(p string)    { configPaths = append(configPaths, p) }

func ReadInConfig() error {
	if len(configPaths) == 0 {
		configPaths = []string{"."}
	}
	path := filepath.Join(configPaths[0], configName+"."+configType)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	return nil
}

func Unmarshal(out interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func WatchConfig()                        {}
func OnConfigChange(func(fsnotify.Event)) {}
