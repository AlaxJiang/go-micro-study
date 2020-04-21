package config

import (
	"fmt"
	"testing"
)

type DataConfig struct {
	DataSourceConfList []DataSourceConf `yaml:"datasources,omitempty"`
	DBAccount          DBAccount        `yaml:"dbData,omitempty"`
}

type DBAccount struct {
	UserName string `yaml:"userName,omitempty"`
	Password string `yaml:"password,omitempty"`
}
type DataSourceConf struct {
	DataSourceLabel string `yaml:"dataSourceLabel,omitempty"`
	MysqlUrl        string `yaml:"mysqlUrl,omitempty"`
}

func TestLoadYml(t *testing.T) {
	var configModel *DataConfig
	configModel = new(DataConfig)
	err := LoadYmlConfig("mysql.yml", configModel)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*configModel)
	}
}

type DataConfigJSON struct {
	DataSourceConfList []DataSourceConfJSON `json:"datasources,omitempty"`
	DBAccount          DBAccountJSON        `json:"dbData,omitempty"`
}

type DBAccountJSON struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}
type DataSourceConfJSON struct {
	DataSourceLabel string `json:"dataSourceLabel,omitempty"`
	MysqlUrl        string `json:"mysqlUrl,omitempty"`
}

func TestLoadJson(t *testing.T) {
	var configModel *DataConfigJSON
	configModel = new(DataConfigJSON)
	err := LoadJsonConfig("mysql.json", configModel)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*configModel)
	}
}
