package models

import (
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
	DBType          string `yaml:"dbType,omitempty"`
	DBUrl           string `yaml:"dbUrl,omitempty"`
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
	DBType          string `json:"dbType,omitempty"`
	DBUrl           string `json:"dbUrl,omitempty"`
}
