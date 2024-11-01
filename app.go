package main

import (
	"context"
	"durability-go/factorial"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
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

// Greet returns a greeting for the given name
func (a *App) GetESLfattoriale(rsl float64, fa float64, fb float64, fc float64, fd float64, fe float64, ff float64, fg float64) float64 {
	fat := factorial.NewFattoriale()
	fat.SetFactors(fa, fb, fc, fd, fe, ff, fg)
	fat.SetRSL(rsl)
	return fat.GetESL()
}
