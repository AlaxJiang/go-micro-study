package models

import (
	"database/sql"
	"fmt"
	"log"
	_ "sync"
)

type ClusterOptions struct {
	dataSourceLabel string
	dbType          string
	dbUrl           string
}

type ClusterOption func(c *ClusterOptions)

func BuilderClusterDbUrl(dbUrl string) ClusterOption {
	return func(opts *ClusterOptions) {
		opts.dbUrl = dbUrl
	}
}
func BuilderClusterDbType(dbType string) ClusterOption {
	return func(opts *ClusterOptions) {
		opts.dbType = dbType
	}
}
func BuilderClusterDataSourceLabel(dataSourceLabel string) ClusterOption {
	return func(opts *ClusterOptions) {
		opts.dataSourceLabel = dataSourceLabel
	}
}

func NewDBCluster(opts ...ClusterOption) (*DBCluster, error) {
	clusterOpts := ClusterOptions{}
	for _, opt := range opts {
		opt(&clusterOpts)
	}
	dbCluster := new(DBCluster)

	//dbCluster.opts = clusterOpts
	//sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", clusterOpts.userName, clusterOpts.passWord, clusterOpts.hostName, clusterOpts.port, clusterOpts.dataBase))
	sqlDB, err := sql.Open(fmt.Sprintf("%s", clusterOpts.dbType), fmt.Sprintf("%s", clusterOpts.dbUrl))
	if err != nil {
		log.Fatal(err.Error())
		return dbCluster, err
	}
	sqlDB.SetConnMaxLifetime(600)
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
		return dbCluster, err
	}
	dbCluster.SqlDB = sqlDB
	return dbCluster, nil
}

/*
var (
	dbConfigMap sync.Map
)

func Init(dataConfig interface{}) (map[string]*DBCluster, DBAccount) {
	fmt.Println("Mysql Init")
	dbClusterMap := make(map[string]*DBCluster)
	for _, defaultMysqlConf := range dataConfig.DataSourceConfList {
		mysqlOpts := []ClusterOption{
			BuilderClusterMysqlurl(defaultMysqlConf.MysqlUrl),
			BuilderClusterDataSourceLabel(defaultMysqlConf.DataSourceLabel),
		}
		if dbCluster, err := NewDBCluster(mysqlOpts...); err == nil {
			dbClusterMap[defaultMysqlConf.DataSourceLabel] = dbCluster
		}
	}
	fmt.Println(dbClusterMap)
	return dbClusterMap, dataConfig.DBAccount
}
*/
