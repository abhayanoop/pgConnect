package pgConnect

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DBHost     string `json:"dbHost"`
	DBPort     int    `json:"dbPort"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBSSLMode  string `json:"dbSSLMode"`
}

func GetPostgresDB(configPath string) (db *sql.DB, err error) {

	viper.SetConfigFile(configPath)

	if err = viper.ReadInConfig(); err != nil {
		err = errors.New("Config file not found in path. Error - " + err.Error())
		return
	}

	var config DBConfig

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword,
		config.DBName, config.DBSSLMode)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return
	}

	return
}
