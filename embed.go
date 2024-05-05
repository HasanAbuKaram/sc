package main

import (
	"embed"
)

//go:embed version.txt
var version string

//go:embed repoUrl.txt
var repoUrl string

//go:embed index.html login.html
var content embed.FS
