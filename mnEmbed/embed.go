package mnEmbed

import (
	"embed"
)

//go:embed template/*
var Templates embed.FS
