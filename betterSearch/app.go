package main

import (
	"betterSearch/backend"
	"context"
	"fmt"
)

// Feels suboptimal that we're putting this here....

// App struct
type App struct {
	ctx          context.Context
	wailsTest    *backend.WailsTest
	dirUtils     *backend.DirUtils
	searchStruct *backend.SearchStruct
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

func (a *App) GenerateDirectoryMap(dir string) []string {
	return a.dirUtils.GenerateDirectoryMap(dir)
}

func (a *App) ChooseDirectory() (string, error) {
	return a.dirUtils.ChooseDirectory()
}

// How do we handle fucking uhhhh structs getting passed between the front and back?
func (a *App) Search(target string) []backend.Result {
	return a.searchStruct.Search(target)
}
