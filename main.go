package main

import (
	"clipr/internal/api"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := api.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "clipr",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WindowIsTranslucent: true,            // or windows.Transparent
			BackdropType:        windows.Acrylic, // optional: Mica, Acrylic, etc.
		},
		BackgroundColour: &options.RGBA{R: 32, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
