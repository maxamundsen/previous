// This package acts as the "global intitialization" package that is called when a process in the codebase starts.
// Preload is especially necessary for cases like compile-time execution, where certain global state must be initialized
package preload

import (
	"net/http"
	"previous/config"
	"previous/pages"
	"previous/repository"
	"previous/security"
	"previous/tasks"
)

// Any data that must be returned to your program from a `preload` call belongs here
type PreloadResourceBundle struct {
	HttpMux *http.ServeMux
}

// Expose options to programs utilizing preload
type PreloadOptions struct {
	ShouldCreateHttpMux bool
	ShouldInitTasks     bool
	ShouldInitDatabase  bool
}

// If you don't care about setting options and just want to include everything, just use this:
func PreloadOptionsAll() PreloadOptions {
	return PreloadOptions{
		ShouldCreateHttpMux: true,
		ShouldInitTasks:     true,
		ShouldInitDatabase:  true,
	}
}

// When adding dependencies to preload, ensure that they are loaded in the correct order.
// For example, database initialization reads from the config, so config must be loaded first.
func Preload(options PreloadOptions) PreloadResourceBundle {
	bundle := PreloadResourceBundle{}

	config.Init()
	security.Init()
	pages.Init()

	if options.ShouldInitDatabase {
		repository.Init()
	}

	if options.ShouldCreateHttpMux {
		bundle.HttpMux = http.NewServeMux()
	}

	if options.ShouldInitTasks {
		tasks.InitTasks()
	}

	return bundle
}
