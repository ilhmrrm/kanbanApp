package public

import (
	"embed"
	_ "embed"
)

//go:embed assets
var AssetsDir embed.FS
