// This package acts as the "global intitialization" package that is called when a process in the codebase starts.
// Preload is especially necessary for cases like compile-time execution, where certain global state must be initialized
package preload

import (
	"net/http"
	"previous/config"
	"previous/repository"
	"previous/security"
	"previous/tasks"
)

// Any data that must be returned to your program from a `preload` call belongs here
var HttpFS  http.Handler
var HttpMux *http.ServeMux

// Expose options to programs utilizing preload
type PreloadOptions struct {
	ShouldCreateHTTPResources bool
	ShouldInitTasks     bool
	ShouldInitDatabase  bool
}

// If you don't care about setting options and just want to include everything, just use this:
func PreloadOptionsAll() PreloadOptions {
	return PreloadOptions{
		ShouldCreateHTTPResources: true,
		ShouldInitTasks:     true,
		ShouldInitDatabase:  true,
	}
}

// When adding dependencies to preload, ensure that they are loaded in the correct order.
// For example, database initialization reads from the config, so config must be loaded first.
func PreloadInit(options PreloadOptions) {
	config.Init()
	security.Init()

	if options.ShouldInitDatabase {
		repository.Init()
	}

	if options.ShouldCreateHTTPResources {
		HttpMux = http.NewServeMux()
		HttpFS = http.FileServer(http.Dir("wwwroot"))
	}

	if options.ShouldInitTasks {
		tasks.InitTasks()
	}
}
