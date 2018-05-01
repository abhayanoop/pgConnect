# pgConnect
Connect to Postgres Database using JSON configuration file

Postgres driver used: github.com/lib/pq
Configuration go package used: github.com/spf13/viper

This code requires a JSON configuration file with database credentials.

Example JSON:

filename - dbconfig.json

{	
	"dbHost": "localhost",
	"dbPort": 5432,
	"dbUser": "postgres",
	"dbPassword": "password",
	"dbName": "my_database_name",
	"dbSSLMode": "disable"
}

The complete path with extension of this JSON file will be the configPath argument

Example function call: 

db, err := pgConnect.GetPostgresDB("/home/admin/dbconfig.json")
if err != nil {
  <Handle error>
}