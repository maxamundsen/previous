package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type metagenConfiguration struct {
	Debug bool `json:"Debug"`
	DbSource string `json:"DbSource"`
	DbConnectionString string `json:"DbConnectionString"`
	DbSchema string `json:"DbSchema"`
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
	"Debug": true,
	"DbSource": "sqlite",
	"DbConnectionString": "file:example.db?parseTime=true",
	"DbSchema": "example"
}`