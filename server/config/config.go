package config

import (
	"log"
	"os"

	"github.com/prynnekey/go-reggie/server/global"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// // 获取全部文件内容
	// fmt.Printf("viper.AllSettings(): %v\n", viper.AllSettings())
	// fmt.Println("--------------")
	//
	// fmt.Printf("viper.GetString(\"database.mysql.datasource.host\"): %v\n", viper.GetString("database.mysql.datasource.host"))
	// fmt.Printf("viper.GetInt(\"database.mysql.datasource.port\"): %v\n", viper.GetInt("database.mysql.datasource.port"))

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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

	global.GVA_DB = db
}
