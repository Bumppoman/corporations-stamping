package main

import(
	"github.com/wailsapp/wails/v3/pkg/application"
)

// App struct
type SettingsService struct {}

// OpenSettings opens the preference pane
func (p *SettingsService) OpenSettings() *application.WebviewWindow {
	app := application.Get()

	// Create a new window with the necessary options
	settingsWindow := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Preferences",
		URL: "/settings",
	})

	return settingsWindow
}

// Set sets a setting
func (p *SettingsService) Set(key string, value string) error {
	return nil
}
