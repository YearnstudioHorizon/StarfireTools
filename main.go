package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	didRelaunch, err := maybeRelaunchAsSystem()
	if err != nil {
		log.Printf("elevation failed: %v", err)
	}
	if didRelaunch {
		return
	}

	// Create an instance of the app structure
	app := NewApp()

	// 获取当前用户名，用于隔离不同权限或用户的 WebView2 数据目录
	// 解决以 SYSTEM 权限运行时 WebView2 在 systemprofile 目录下没有权限读写的问题
	userName := os.Getenv("USERNAME")
	if userName == "" {
		userName = "Default"
	}
	userDataDir := filepath.Join(os.Getenv("PROGRAMDATA"), "StarFireTool", "WebView2Data_"+userName)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "StarfireTool",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []any{
			app,
		},
		Frameless: true,
		Windows: &windows.Options{
			WebviewUserDataPath: userDataDir,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
