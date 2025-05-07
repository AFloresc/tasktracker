package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	DBPort     string
	DBUser     string
	DBDSN      string
	DBPassword string
	DBHostName string
	DBName     string
}

func GetConfiguration() Configuration {
	var c Configuration

	c.DBPort = getenv("PORT", ":8080")
	c.DBUser = getenv("DBUSER", "root")
	c.DBName = getenv("MYSQL_DBNAME", "db_tasktracker")
	c.DBPassword = getenv("MYSQL_PASSWORD", "12345.")
	c.DBHostName = getenv("MYSQL_HOSTNAME", "localhost")

	c.DBDSN = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", getenv("MYSQL_USER", "root"), getenv("MYSQL_PASSWORD", "12345."), getenv("MYSQL_HOSTNAME", "localhost"), getenv("MYSQL_DBNAME", "db_tasktracker"))

	return c
}

func getenv(key, fallback string) string {
	// log.Println(os.Getenv(key))
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
