// This package acts as the "global intitialization" package that is called when a process in the codebase starts.
// Preload is especially necessary for cases like compile-time execution, where certain global state must be initialized
package preload

import (
	"net/http"
	"previous/config"
	"previous/pages"
	"previous/repository"
	"previous/tasks"
)

type PreloadResourceBundle struct {
	HttpMux *http.ServeMux
}

type PreloadOptions struct {
	ShouldCreateHttpMux bool
	ShouldInitTasks     bool
	ShouldInitDatabase  bool
}

func PreloadOptionsAll() PreloadOptions {
	return PreloadOptions{
		ShouldCreateHttpMux: true,
		ShouldInitTasks: true,
		ShouldInitDatabase: true,
	}
}

func Preload(options PreloadOptions) PreloadResourceBundle {
	bundle := PreloadResourceBundle{}

	config.LoadConfig()
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
