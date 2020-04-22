package configs

import (
	"errors"
	"path/filepath"
	"strings"

	"os"

	_ "github.com/go-yaml/yaml"
	"github.com/micro/go-micro/config"
	"gopkg.in/yaml.v3"
)

var (
	ymlFileExtension string = ".yml"
)

func LoadYmlConfig(configFilePath string, configModel interface{}) error {
	if !strings.HasSuffix(configFilePath, ymlFileExtension) {
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
	err = yaml.Unmarshal(config.Bytes(), configModel)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}
