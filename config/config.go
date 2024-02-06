package config

import (
	"encoding/json"
	"webdawgengine/build"
	"io"
	"log"
	"os"
	"reflect"
)

// the config package allows the use of runtime configuration options from a file.
// the config options live in a global, non-exported struct called 'config'.
// the InitConfiguration() function reads data from a config file, and
// populates this struct. This function should be called once when the program
// starts.

// typically the configuration is read from `config.json`, however
// when the `devel` build tag is set, the options are read from `config.devel.json`

// the config struct is not exported, and must be retrieved via the GetConfiguration() function.
// because of this, you cannot modify the values inside configuration from anywhere in the program
// (except the `config` package directly)

// define all configuration options here
type configuration struct {
	Host             string `json:"Host"`
	CookieExpiryDays string `json:"CookieExpiryDays"`
	ConnectionString string `json:"ConnectionString"`
	SmtpServer       string `json:"SmtpServer"`
	SmtpPort         string `json:"SmtpPort"`
	SmtpUsername     string `json:"SmtpUsername"`
	SmtpDisplayFrom  string `json:"SmtpDisplayFrom"`
	SmtpPassword     string `json:"SmtpPassword"`
	SmtpRequireAuth  string `json:"SmtpRequireAuth"`
}

var config configuration


// log warning if provided struct member is not present in json config
func warnMissingMember(member interface{}) {
	log.Printf("Configuration warning: missing option: `%s`, type: `%s`.", getConfigFieldName(member), reflect.TypeOf(member))
}

// pretty terrible way to check, but this language has poor meta-programming ¯\_(ツ)_/¯
func getConfigFieldName(member interface{}) string {
    structType := reflect.TypeOf(config)
    for i := 0; i < structType.NumField(); i++ {
        field := structType.Field(i)
        fieldValue := reflect.ValueOf(config).Field(i).Interface()
        if reflect.DeepEqual(fieldValue, member) {
            return field.Name
        }
    }
    return ""
}

// specify default values if none are specified in the config file
func setDefaultValues() {
	if config.Host == "" {
		warnMissingMember(config.Host)
		config.Host = "localhost:8080"
	}

	if config.CookieExpiryDays == "" {
		warnMissingMember(config.CookieExpiryDays)
		config.CookieExpiryDays = "7"
	}

	if config.ConnectionString == "" {
		warnMissingMember(config.ConnectionString)
	}

	if config.SmtpServer == "" {
		warnMissingMember(config.SmtpServer)
	}

	if config.SmtpPort == "" {
		warnMissingMember(config.SmtpPort)
	}

	if config.SmtpUsername == "" {
		warnMissingMember(config.SmtpUsername)
	}

	if config.SmtpDisplayFrom == "" {
		warnMissingMember(config.SmtpDisplayFrom)
	}

	if config.SmtpPassword == "" {
		warnMissingMember(config.SmtpPassword)
	}

	if config.SmtpRequireAuth == "" {
		warnMissingMember(config.SmtpRequireAuth)
		config.SmtpRequireAuth = "true"
	}
}


func ReadConfiguration() {
	var configFile *os.File
	var err error

	if build.DEVEL {
		// if a config.devel.json exists
		configFile, err = os.Open("config.devel.json")

		// fallback to default config if devel config does not exist
		if err != nil {
			configFile, err = os.Open("config.json")
		}
	} else {
		configFile, err = os.Open("config.json")
	}

	if err != nil {
		log.Fatal(err)
	}

	configBytes, readErr := io.ReadAll(configFile)

	if readErr != nil {
		log.Fatal(err)
	}

	defer configFile.Close()

	json.Unmarshal(configBytes, &config)

	setDefaultValues()
	log.Println("Loaded configuration file [" + configFile.Name() + "]")
}

// return a COPY of the configuration to ensure that
// config options cannot be modified globally
func GetConfiguration() configuration {
	return config
}
