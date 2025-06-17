package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"fileExploreGo/goFiles"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	foo := NewUtils()
	files := goFiles.NewFiles()
	myFuzzySearch := goFiles.NewFuzzySearch()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "fileExploreGo",
		Width:  1000,
		Height: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			foo,
			files,
			myFuzzySearch,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
