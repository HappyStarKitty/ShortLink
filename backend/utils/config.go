// 没有用到
package utils

import "github.com/spf13/viper"

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetServerAddress() string {
	return viper.GetString("server.address")
}

func GetDBHost() string {
	return viper.GetString("db.host")
}

func GetDBPort() int {
	return viper.GetInt("db.port")
}

func GetDBUser() string {
	return viper.GetString("db.user")
}

func GetDBName() string {
	return viper.GetString("db.name")
}

func GetDBPassword() string {
	return viper.GetString("db.password")
}
