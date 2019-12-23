package bootstrap

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"net/url"
	"newland/app/libraries/file"
)

var MasterDataSourceName string
var SlaveDataSourceName string

var MysqlEngine = make(map[string]*xorm.EngineGroup)
var Timezone = url.QueryEscape("Asia/shanghai")


type MysqlDns struct {
	Host	 string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type MysqlConnectGroup struct {
	Name   string     `yaml:"name"`
	Master MysqlDns   `yaml:"master"`
	Slave  []MysqlDns `yaml:"slave"`
}

func MysqlInit() {
	dbConfigStr := file.GetFileContent("/Users/banxuelai/go/src/newland/config/test/db.yml")
	mysqlConfigList := make([]MysqlConnectGroup,0)
	err := yaml.Unmarshal(dbConfigStr, &mysqlConfigList)
	if err != nil {
		panic(err)
	}
	for _, mysqlConfigItem := range mysqlConfigList {
		MasterDataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true", mysqlConfigItem.Master.UserName, mysqlConfigItem.Master.Password, mysqlConfigItem.Master.Host, mysqlConfigItem.Master.Port, mysqlConfigItem.Master.Database, Timezone)
		master, err := xorm.NewEngine("mysql", MasterDataSourceName)
		//最大连接数
		master.SetMaxIdleConns(200)
		//空闲连接
		master.SetMaxOpenConns(50)
		if err != nil {
			panic(err)
		}
		slaves := []*xorm.Engine{}

		for _, s := range mysqlConfigItem.Slave {
			SlaveDataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true", s.UserName, s.Password, s.Host, s.Port, s.Database, Timezone)
			salve, err := xorm.NewEngine("mysql", SlaveDataSourceName)
			salve.SetMaxIdleConns(200)
			salve.SetMaxOpenConns(50)
			if err != nil {
				panic(err)
			}
			slaves = append(slaves, salve)
		}
		MysqlEngine[mysqlConfigItem.Name], err = xorm.NewEngineGroup(master, slaves)

		if err != nil {
			panic(err)
		}
	}
}