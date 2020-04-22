package configs

import (
	"fmt"
	"testing"

	"./go-micro-study/models"
)

func TestLoadYml(t *testing.T) {

	var configModel *models.DataConfig
	configModel = new(DataConfig)
	err := LoadYmlConfig("mysql.yml", configModel)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*configModel)
	}
}

func TestLoadJson(t *testing.T) {
	var configModel *models.DataConfigJSON
	configModel = new(DataConfigJSON)
	err := LoadJsonConfig("mysql.json", configModel)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*configModel)
	}
}
