package config

import (
	"log"

	"github.com/spf13/viper"
)

type SqliteConfig struct {
	Dbc string
	Db string
}

func GetSqliteConfig() *SqliteConfig {
	
	viper.SetConfigFile("/usr/local/src/go/go-orm/config/config.yaml")

	if err := viper.ReadInConfig() ; err != nil {
		log.Fatalln(err)
	}
	
	return &SqliteConfig{
		Dbc: viper.GetString("sqlite.name"),
		Db: viper.GetString("sqlite.db"),
	}
}