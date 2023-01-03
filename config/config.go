package config

import (
	"log"
	"os"

	"github.com/prynnekey/go-reggie/global"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() {
	// 获取当前项目目录
	dir, err := os.Getwd()
	if err != nil {
		log.Println("os.Getwd err:", err)
	}
	// 设置文件名和文件后缀
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// 配置文件所在文件夹
	viper.AddConfigPath(dir + "/config")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	initMysql()

}

func initMysql() {
	gormMysql := Mysql{
		Host:     viper.GetString("database.mysql.datasource.host"),
		port:     viper.GetString("database.mysql.datasource.port"),
		Username: viper.GetString("database.mysql.datasource.username"),
		Password: viper.GetString("database.mysql.datasource.password"),
		Dbname:   viper.GetString("database.mysql.datasource.dbname"),
	}
	dsn := gormMysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 开启显示sql日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	d, err2 := db.DB()
	if err2 != nil {
		panic(err2)
	}
	err3 := d.Ping()
	if err3 != nil {
		panic(err3)
	}

	// fmt.Println("连接数据库成功")

	global.DB = db
}
