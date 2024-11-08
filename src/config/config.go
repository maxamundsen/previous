// The config package allows an easy way to provide variable options to the
// program read at runtime during startup. These are typically settings like the
// http port, database connection string, or password.
package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	SESSION_COOKIE_NAME            = "_webdawgengine_session"
	SESSION_COOKIE_EXPIRY_DAYS int = 100
	SESSION_COOKIE_ENTROPY     int = 33

	IDENTITY_COOKIE_NAME        string = "_webdawgengine_identity"
	IDENTITY_COOKIE_EXPIRY_DAYS int    = 30
	IDENTITY_TOKEN_EXPIRY_DAYS  int    = 30
	IDENTITY_COOKIE_ENTROPY     int    = 33
	IDENTITY_LOGIN_PATH         string = "/auth/login"
	IDENTITY_LOGOUT_PATH        string = "/auth/logout"
	IDENTITY_DEFAULT_PATH       string = "/app/dashboard"
	IDENTITY_AUTH_REDIRECT      bool   = true

	PASSWORD_MIN_LENGTH         int = 8
	PASSWORD_REQUIRED_UPPERCASE int = 1
	PASSWORD_REQUIRED_LOWERCASE int = 1
	PASSWORD_REQUIRED_NUMBERS   int = 1
	PASSWORD_REQUIRED_SYMBOLS   int = 0

	MAX_LOGIN_ATTEMPTS int = 5
)

type configuration struct {
	Host                    string `json:"Host"`
	Port                    string `json:"Port"`
	IdentityPrivateKey      string `json:"IdentityPrivateKey"`
	IdentityDefaultPassword string `json:"IdentityDefaultPassword"`
	SessionPrivateKey       string `json:"SessionPrivateKey"`
	DbConnectionString      string `json:"DbConnectionString"`
	SmtpServer              string `json:"SmtpServer"`
	SmtpPort                string `json:"SmtpPort"`
	SmtpUsername            string `json:"SmtpUsername"`
	SmtpDisplayFrom         string `json:"SmtpDisplayFrom"`
	SmtpPassword            string `json:"SmtpPassword"`
	SmtpRequireAuth         bool   `json:"SmtpRequireAuth"`
}

var config configuration

func GetConfig() configuration {
	return config
}

func LoadConfig(configFile *os.File) {
	configBytes, readErr := io.ReadAll(configFile)

	if readErr != nil {
		log.Fatal(readErr)
	}

	defer configFile.Close()

	jsonErr := json.Unmarshal(configBytes, &config)

	if jsonErr != nil {
		log.Fatal(fmt.Errorf("configuration error: %v", jsonErr))
	}

	log.Println("Loaded configuration file [" + configFile.Name() + "]")
}
