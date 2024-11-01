package factorial

type Fattoriale struct {
	A, B, C, D, E, F, G, RSL float64
}

func NewFattoriale() *Fattoriale {
	return &Fattoriale{
		A: 1.0,
		B: 1.0,
		C: 1.0,
		D: 1.0,
		E: 1.0,
		F: 1.0,
		G: 1.0,
	}
}

func (f *Fattoriale) SetA(a float64) {
	f.A = a
}

func (f *Fattoriale) SetB(b float64) {
	f.B = b
}

func (f *Fattoriale) SetC(c float64) {
	f.C = c
}

func (f *Fattoriale) SetD(d float64) {
	f.D = d
}

func (f *Fattoriale) SetE(e float64) {
	f.E = e
}

func (f *Fattoriale) SetF(g float64) {
	f.F = g
}

func (f *Fattoriale) SetG(g float64) {
	f.G = g
}

func (f *Fattoriale) SetRSL(rsl float64) {
	f.RSL = rsl
}

func (f *Fattoriale) SetFactors(a, b, c, d, e, fVal, g float64) {
	f.A = a
	f.B = b
	f.C = c
	f.D = d
	f.E = e
	f.F = fVal
	f.G = g
}

func (f *Fattoriale) GetA() float64 {
	return f.A
}

func (f *Fattoriale) GetB() float64 {
	return f.B
}

func (f *Fattoriale) GetC() float64 {
	return f.C
}

func (f *Fattoriale) GetD() float64 {
	return f.D
}

func (f *Fattoriale) GetE() float64 {
	return f.E
}

func (f *Fattoriale) GetF() float64 {
	return f.F
}

func (f *Fattoriale) GetG() float64 {
	return f.G
}

func (f *Fattoriale) GetRSL() float64 {
	return f.RSL
}

func (f *Fattoriale) GetFactors() map[string]float64 {
	return map[string]float64{
		"A": f.A,
		"B": f.B,
		"C": f.C,
		"D": f.D,
		"E": f.E,
		"F": f.F,
		"G": f.G,
	}
}

func (f *Fattoriale) GetESL() float64 {
	return f.RSL * f.A * f.B * f.C * f.D * f.E * f.F * f.G
}
