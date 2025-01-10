package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type metagenConfiguration struct {
	DevDbType string `json:"DevDbType"`
	DevDbConnectionString string `json:"DevDbConnectionString"`
	DevDbSchema string `json:"DevDbSchema"`

	ProdDbType string `json:"ProdDbType"`
	ProdDbConnectionString string `json:"ProdDbConnectionString"`
	ProdDbSchema string `json:"ProdDbSchema"`

	StagingDbType string `json:"StagingDbType"`
	StagingDbConnectionString string `json:"StagingDbConnectionString"`
	StagingDbSchema string `json:"StagingDbSchema"`
}

var config metagenConfiguration

func loadMetaConfig() {
	configFile, err := os.Open("metagen_config.json")

	var configBytes []byte
	var readErr error

	if configFile != nil && err == nil {
		configBytes, readErr = io.ReadAll(configFile)
		defer configFile.Close()

		if readErr != nil {
			log.Fatal(readErr)
		}
	} else {
		//attempt to generate default config
		log.Println("Config file not found, attempting to generate a default.")
		configBytes = []byte(DEFAULT_METAGEN_CONFIG)

		f, err := os.Create("metagen_config.json")
		if err != nil {
			log.Fatal("Error generating default metagen config file. Aborting.")
		}

		defer f.Close()

		_, err = f.Write(configBytes)
		if err != nil {
			log.Fatal("Error writing the default metagen config to disk. Aborting")
		}
	}

	jsonErr := json.Unmarshal(configBytes, &config)

	if jsonErr != nil {
		log.Fatal(fmt.Errorf("Config error: %v", jsonErr))
	}

}

const DEFAULT_METAGEN_CONFIG = `{
	"DevDbType": "sqlite",
	"DevDbConnectionString": "file:example.db",
	"DevDbSchema": "example",

	"ProdDbType": "sqlite",
	"ProdDbConnectionString": "file:example.db",
	"ProdDbSchema": "example",

	"StagingDbType": "sqlite",
	"StagingDbConnectionString": "file:example.db",
	"StagingDbSchema": "example",
}`