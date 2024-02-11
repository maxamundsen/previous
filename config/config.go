package config

import (
	"encoding/json"
	"webdawgengine/build"
	"io"
	"log"
	"os"
	"reflect"
	"fmt"
)

// the config package allows the use of runtime configuration options from a file.
// the config options live in a global, non-exported struct called 'config'.
// the ParseConfigFile() function reads data from a config file, and
// populates this struct. This function should be called once when the program
// starts.

// To use the config, you must call a getter function that has been exported
// since the configuration struct is 'private' to the config package

// If the config file is not valid JSON, the program will throw an error, and exit.
// If you are missing any value from the config file, the missing value will
// be printed to the console, with the expected type.

// typically the configuration is read from `config.json`, however
// when the `devel` build tag is set, the options are read from `config.devel.json`


// define all configuration options here
type configuration struct {
	Host             *string `json:"Host"`
	CookieExpiryDays *int    `json:"CookieExpiryDays"`
	ConnectionString *string `json:"ConnectionString"`
	SmtpServer       *string `json:"SmtpServer"`
	SmtpPort         *string `json:"SmtpPort"`
	SmtpUsername     *string `json:"SmtpUsername"`
	SmtpDisplayFrom  *string `json:"SmtpDisplayFrom"`
	SmtpPassword     *string `json:"SmtpPassword"`
	SmtpRequireAuth  *bool   `json:"SmtpRequireAuth"`
}

var config configuration

// getter functions
func GetHost() string {
	return *config.Host
}

func GetCookieExpiryDays() int {
	return *config.CookieExpiryDays
}

func GetConnectionString() string {
	return *config.ConnectionString
}

func GetSmtpServer() string {
	return *config.SmtpServer
}

func GetSmtpPort() string {
	return *config.SmtpPort
}

func GetSmtpUsername() string {
	return *config.SmtpUsername
}

func GetSmtpDisplayFrom() string {
	return *config.SmtpDisplayFrom
}

func GetSmtpPassword() string {
	return *config.SmtpPassword
}

func GetSmtpRequireAuth() bool {
	return *config.SmtpRequireAuth
}

// when the config file is parsed, each field in the config struct is checked for nil.
// any field that is nil is logged as such to the console using reflection
// "metaprogramming" :D
func checkMissingFields(parsed configuration, expected configuration) {
	expectedType := reflect.TypeOf(expected)

	for i := 0; i < expectedType.NumField(); i++ {
		field := expectedType.Field(i)
		fieldType := field.Type.Elem()

		// Check if the field is the zero value for its type in the parsed struct
		if reflect.ValueOf(parsed).Field(i).IsNil() {
			log.Fatalf("Configuration error: missing option: `%s`, expecting type: `%s`", field.Tag.Get("json"), fieldType)
		}
	}
}

func ParseConfigFile() {
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

	jsonErr := json.Unmarshal(configBytes, &config)

	if jsonErr != nil {
		log.Fatal(fmt.Errorf("Configuration error: %v", jsonErr))
	}

	checkMissingFields(config, configuration{})

	log.Println("Loaded configuration file [" + configFile.Name() + "]")
}