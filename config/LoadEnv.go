package config

import (
	"fmt"
	"os"

	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"github.com/gobuffalo/packr/v2"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file

func LoadEnvVar() {
	box := packr.New("config", "../env")

	str,err := box.FindString(".env")
    fmt.Println("🚀 ", str)
	
	logs.ERROR("Error loading .env file",err)

	// load .env file
	allEnvVar ,err := godotenv.Unmarshal(str)
	logs.ERROR("Error unmarshalling in godotenv.Unmarshal(str) ",err)

	for i,j := range allEnvVar{
		os.Setenv(i,j)
	}
}