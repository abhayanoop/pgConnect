package pgConnect

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Go structure to unmarshal JSON config file

type DBConfig struct {
	DBHost     string `json:"dbHost"`
	DBPort     int    `json:"dbPort"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBSSLMode  string `json:"dbSSLMode"`
}

// Opens a postgres DB connection with credentials as given in a JSON config file
// The paths in order of priority and name of config file (without extension) should be given as argument.
// Refer README.md for example JSON

func GetPostgresDB(configPaths []string, configName string) (db *sql.DB, err error) {

	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName(configName)

	if err = viper.ReadInConfig(); err != nil {
		err = errors.New("Config file not found in path. Error - " + err.Error())
		return
	}

	var config DBConfig

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	pgInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword,
		config.DBName, config.DBSSLMode)

	if db, err = sql.Open("postgres", pgInfo); err != nil {
		return
	}

	return
}
