package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configurations struct {
	Server         ServerConfigurations
	Database       DatabaseConfigurations
	ProductionEnv  string
	DevelopmentEnv string
	EnvVariable    string
	CSVName        string
}

type ServerConfigurations struct {
	Port string
}

type DatabaseConfigurations struct {
	ProductionDBName  string
	DevelopmentDBName string
}

func (c *Configurations) SetConfig() {
	c.ProductionEnv = "./production.env"
	c.DevelopmentEnv = "../../../development.env"
	c.Server.Port = ":3000"
	c.Database.ProductionDBName = "storage.db"
	c.Database.DevelopmentDBName = "../../../test.db"
	c.EnvVariable = "WEATHER_API_KEY"
	c.CSVName = "tasks.csv"
}

func LoadEnv(s string, c Configurations) string {
	err := godotenv.Load(s)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	keys := os.Getenv(c.EnvVariable)

	return keys
}
