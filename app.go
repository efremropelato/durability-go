package main

import (
	"context"
	"durability-go/factorial"
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

func (a *App) GetESLSimple(rsl float64, fa float64, fb float64, fc float64, fd float64, fe float64, ff float64, fg float64) float64 {
	fat := factorial.NewFattoriale()
	fat.SetFactors(fa, fb, fc, fd, fe, ff, fg)
	fat.SetRSL(rsl)
	return fat.GetESL()
}

func (a *App) GetESLEvolute(rsl float64, iterations int,
	fat int, fa1 float64, fa2 float64, fa3 float64,
	fbt int, fb1 float64, fb2 float64, fb3 float64,
	fct int, fc1 float64, fc2 float64, fc3 float64,
	fdt int, fd1 float64, fd2 float64, fd3 float64,
	fet int, fe1 float64, fe2 float64, fe3 float64,
	fft int, ff1 float64, ff2 float64, ff3 float64,
	fgt int, fg1 float64, fg2 float64, fg3 float64) []byte {
	evo := factorial.Evolute{}
	evo.SetN(iterations)
	evo.SetM(1)
	evo.SetA(fat, fa1, fa2, fa3, iterations, 1)
	evo.SetB(fbt, fb1, fb2, fb3, iterations, 1)
	evo.SetC(fct, fc1, fc2, fc3, iterations, 1)
	evo.SetD(fdt, fd1, fd2, fd3, iterations, 1)
	evo.SetE(fet, fe1, fe2, fe3, iterations, 1)
	evo.SetF(fft, ff1, ff2, ff3, iterations, 1)
	evo.SetG(fgt, fg1, fg2, fg3, iterations, 1)
	evo.SetRSL(rsl)

	evo.SetESL()
	return evo.GetEvoluteJson()
}
