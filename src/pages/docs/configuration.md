# Global Configuration

In WDE, the configuration is a global structure that lives in the `config` package.
You can read the contents of the configuration struct using the `config.GetConfig()` function.

```go
type configuration struct {
	Host string `json:"Host"`
	Port string `json:"Port"`
	...
}

var config configuration

func GetConfig() configuration {
	return config
}
```

To avoid any potential mishaps, the configuration struct can only be modified within the `config` package.
`config.GetConfig()` returns a _copy_ of a `configuration` struct, to avoid overwriting the contents of the struct.

Example usage:
```go
package yourpackage

import "webdawgengine/config"

func printHostAddress() {
	println(config.GetConfig().Host)
}
```
## Runtime Configuration
Many web applications require specific runtime settings.
You may use this for configuring a production environment differently than a local development environment, for instance.

WDE ships with a configuration parser that reads from `config.json`, found in the `/src` directory of the project.
When you first download the project template, you will not find a `config.json` file.
This is because only a `config.template.json` file is stored in git.
`config.json` is _specifically added to the `.gitignore` file to avoid leaking secrets in source control_.

```json
{
    "Host": "localhost",
    "Port": "8080",
    "SessionPrivateKey": "key",
    "IdentityPrivateKey": "key",
    "IdentityDefaultPassword": "password"
    "DbConnectionString": "root:PASSWORD@tcp(localhost:3306)/hotlap?parseTime=true",
    "SmtpServer": "server",
    "SmtpPort": "587",
    "SmtpUsername": "username",
    "SmtpDisplayFrom": "displayname",
    "SmtpPassword": "password",
    "SmtpRequireAuth": true,
}
```

## Compile-time Configuration

```go
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
```