package factorial

type Evolute struct {
	A, B, C, D, E, F, G *Factor
	RSL                 float64
}

var N, M int

func (e *Evolute) SetN(n int) {
	N = n
}

func (e *Evolute) SetM(m int) {
	M = m
}

func (e *Evolute) SetA(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.A = &Factor{}
	e.A.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetB(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.B = &Factor{}
	e.B.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetC(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.C = &Factor{}
	e.C.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetD(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.D = &Factor{}
	e.D.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetE(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.E = &Factor{}
	e.E.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetF(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.F = &Factor{}
	e.F.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetG(Type int, arg1, arg2, arg3 float64, N, M int) {
	e.G = &Factor{}
	e.G.SetFactor(Type, arg1, arg2, arg3, N, M)
}

func (e *Evolute) SetRSL(rsl float64) {
	e.RSL = rsl
}

func (e *Evolute) GetA() *Factor {
	return e.A
}

func (e *Evolute) GetB() *Factor {
	return e.B
}

func (e *Evolute) GetC() *Factor {
	return e.C
}

func (e *Evolute) GetD() *Factor {
	return e.D
}

func (e *Evolute) GetE() *Factor {
	return e.E
}

func (e *Evolute) GetF() *Factor {
	return e.F
}

func (e *Evolute) GetG() *Factor {
	return e.G
}

func (e *Evolute) GetRSL() float64 {
	return e.RSL
}

type EvolutaESL struct {
	Mean              float64
	Max               float64
	Min               float64
	Median            float64
	Variance          float64
	StandardDeviation float64
	Skewness          float64
	Values            []float64
}

func (e *EvolutaESL) GetESL(facA, facB, facC, facD, facE, facF, facG []float64, RSL float64, N int) {
	values := make([]float64, N)
	for in := 0; in < N; in++ {
		values[in] = RSL * facA[in] * facB[in] * facC[in] * facD[in] * facE[in] * facF[in] * facG[in]
	}
	e.Values = values

	// Calcolo delle statistiche
	e.Mean = Mean(e.Values)
	e.Max = Max(e.Values)
	e.Min = Min(e.Values)
	e.Median = Median(e.Values)
	e.Variance = Variance(e.Values)
	e.StandardDeviation = StdDev(e.Values)
	e.Skewness = Skewness(e.Values)

}
