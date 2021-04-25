package embed

import (
	"embed"
)

//go:embed template/*
var Templates embed.FS
