# Global Configuration

"Configuration" in WDE is loosely defined as a set of global variables, and constants that can be accessed from anywhere, to set certain settings within the WDE server.
For example, you want to specify the maximum number of login attempts, or set the global database credentials.
Really anything that could be a read-only global variable constitutes itself as a "configuration option".


Out of the box, WDE comes with all of the necessary configuration options needed to configure the included features / examples.
You can add your own options, or remove ones you don't need.


## Runtime Configuration

Runtime configuration options are designed to be set by the system administrator deploying the WDE server executable.
Setting these options does _not_ require rebuilding the application, or access to the source code, as they are set at runtime.

Example: You want to set different database credentials on your local machine, than on the production server.

The runtime config is defined as a struct of type `configuration` in the `config` package shown below.
A private instance of this struct is defined at the top level of the package, and can be accessed from other packages using the `config.GetConfig()` function.
The runtime configuration is retrived using a getter function to prevent the consumer from accidentally overriding its members.

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

### Reading the config file

In order to populate the `config` struct, `config.LoadConfig()` is called from `main()`.

This function attempts to parse `config.json` located in the `/src` directory of the project upon program start.
If the config file is not found, one will be generated for you.
JSON keys must match the struct tags found on the `configuration` struct definition.

The config file looks something like this:
```json
{
    "Host": "localhost",
    "Port": "9090",
	...
}
```

_Note: `config.json` is deliberately added to `.gitignore` to avoid leaking secrets in source control_.

## Compile-time Configuration

Unlike the runtime options, compile-time options are specifically intended to be set by the _developer_ (that's you!).

These are just plain global constants.
They are intentially located in the `config` package (and not elsewhere) to keep all configuration related constants in one place.

```go
const (
	SESSION_COOKIE_NAME            = "_webdawgengine_session"
	SESSION_COOKIE_EXPIRY_DAYS int = 100
	SESSION_COOKIE_ENTROPY     int = 33

	...
)
```


## Example usage
To use the config, import `webdawgengine/config` from any package to get your config options.

```go
package yourpackage

import "webdawgengine/config"

func connectToDatabase() {
	// runtime config
	db, err := sql.Connect("dbengine", config.GetConfig().DbConnectionString)
	...
}

func printSessionName() {
	// compile-time constant config
	println(config.SESSION_COOKIE_NAME)
}
```