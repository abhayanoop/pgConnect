# pgConnect
Connect to Postgres Database using JSON configuration file    

Postgres driver used: &nbsp;  [github.com/lib/pq](github.com/lib/pq)    
Configuration go package used: &nbsp;  [github.com/spf13/viper](github.com/spf13/viper)

This code requires a JSON configuration file with database credentials.  

Example JSON:    

filename - dbconfig.json    

```json
{	  
  "dbHost": "localhost",  
  "dbPort": 5432,  
  "dbUser": "postgres",  
  "dbPassword": "password",  
  "dbName": "my_database_name",  
  "dbSSLMode": "disable"  
}
```         

The path(s) of this JSON file as a slice is the configPaths argument.
Multiple config paths can be provided priority based in the slice.
The filename without extension is the configName argument.

Example function call:   
  
```
db, err := pgConnect.GetPostgresDB([]string{".", $HOME}, "config")
if err != nil {
  <Handle error>
}  

```