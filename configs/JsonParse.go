package configs

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro/go-micro/config"
)

var (
	jsonFileExtension string = ".json"
)

func LoadJsonConfig(configFilePath string, dataConfig interface{}) error {
	if !strings.HasSuffix(configFilePath, jsonFileExtension) {
		return errors.New("config file extension err:" + "文件格式不正确")
	}
	_, err := os.Stat(configFilePath)
	if err != nil {
		configFilePath = "." + string(filepath.Separator) + configFilePath
	}
	err = config.LoadFile(configFilePath)
	if err != nil {
		return errors.New("Could not load config file: %s" + err.Error())
	}
	err = json.Unmarshal(config.Bytes(), dataConfig)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}
