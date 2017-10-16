package swaggen

import "github.com/goline/lapi"

type AppSource struct {
	App lapi.App
}

type SwaggerSource struct {
	// Type represents for swagger file's extension, e.g, yml, json
	Type string

	// Path represents for path to file
	Path string
}

type FolderSource struct {
	// Path represents for path to directory
	// Spec will scan this folder for appropriate comments
	Path string
}

type FileSource struct {
	// Path represents for path to file
	// Spec will scan this file for appropriate comments
	Path string
}
