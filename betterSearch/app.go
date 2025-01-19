package main

import (
	"betterSearch/backend"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx       context.Context
	wailsTest *backend.WailsTest
	dirUtils  *backend.DirUtils
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GoHello() string {
	return a.wailsTest.GoHello()
}

func (a *App) GetDirectoryMap(dir string) []string {
	return a.GetDirectoryMap(dir)
}

func (a *App) openFileInOS(dir string) error {
	return a.openFileInOS(dir)
}
