package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strings"
	"todo_list/model"
)

var (
	AppMode string
	HttpPort string
	RedisAddr string
	RedisPwd string
	RedisDBName string
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPwd string
	DbName string
)

func Init()  {
	file, err := ini.Load("./conf/config.ini");if err != nil {
		fmt.Println("配置文件读取失败", err)
		os.Exit(1)
	}
	LoadService(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPwd, "@tcp(",DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	model.DataBase(path)
}

func LoadService(file *ini.File)  {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File)  {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPwd = file.Section("mysql").Key("DbPwd").String()
	DbName = file.Section("mysql").Key("DbName").String()
}