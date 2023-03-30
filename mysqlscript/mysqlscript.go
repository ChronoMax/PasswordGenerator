package mysqlscript

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username     string `json:"username"`
	Protocol     string `json:"protocol`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"databaseName"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func InsertToDatabase(result string) {
	config, _ := LoadConfiguration("./config.json")

	//db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test")
	var connection string = "" + config.Username + "@" + config.Protocol + "(" + config.Host + ":" + config.Port + ")" + config.DatabaseName + ""
	db, err := sql.Open("mysql", connection)
	if err != nil {
		fmt.Println("Cannot find MSQL database!")
	}
	defer db.Close()

	// prepare a statement for inserting data
	stmt, err := db.Prepare("INSERT INTO passwords(passwords) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// execute the statement with parameter values
	stmt.Exec(result)
}
