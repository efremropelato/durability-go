package factorial

import "encoding/json"

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

type Evolute struct {
	A, B, C, D, E, F, G *Factor
	RSL                 float64
	ESL                 *EvolutaESL
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

func (e *Evolute) SetESL() {
	e.ESL = &EvolutaESL{}
	facA := e.GetA().GetValues(0)
	facB := e.GetB().GetValues(0)
	facC := e.GetC().GetValues(0)
	facD := e.GetD().GetValues(0)
	facE := e.GetE().GetValues(0)
	facF := e.GetF().GetValues(0)
	facG := e.GetG().GetValues(0)

	values := make([]float64, N)
	for in := 0; in < N; in++ {
		values[in] = e.RSL * facA[in] * facB[in] * facC[in] * facD[in] * facE[in] * facF[in] * facG[in]
	}
	e.ESL.Values = values

	e.ESL.Mean = Mean(e.ESL.Values)
	e.ESL.Max = Max(e.ESL.Values)
	e.ESL.Min = Min(e.ESL.Values)
	e.ESL.Median = Median(e.ESL.Values)
	e.ESL.Variance = Variance(e.ESL.Values)
	e.ESL.StandardDeviation = StdDev(e.ESL.Values)
	e.ESL.Skewness = Skewness(e.ESL.Values)
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

func (e *Evolute) GetESL() *EvolutaESL {
	return e.ESL
}

func (e *Evolute) GetEvolute() *Evolute {
	return e
}

type EvoluteResp struct {
	FactorA Factor     `json:"a"`
	FactorB Factor     `json:"b"`
	FactorC Factor     `json:"c"`
	FactorD Factor     `json:"d"`
	FactorE Factor     `json:"e"`
	FactorF Factor     `json:"f"`
	FactorG Factor     `json:"g"`
	RSL     float64    `json:"rsl"`
	ESL     EvolutaESL `json:"esl"`
}

func (e *Evolute) GetEvoluteJson() []byte {

	dd := EvoluteResp{}
	dd.FactorA = *e.GetA()
	dd.FactorB = *e.GetB()
	dd.FactorC = *e.GetC()
	dd.FactorD = *e.GetD()
	dd.FactorE = *e.GetE()
	dd.FactorF = *e.GetF()
	dd.FactorG = *e.GetG()
	dd.RSL = e.GetRSL()
	dd.ESL = *e.GetESL()
	jsonResp, _ := json.Marshal(dd)

	return jsonResp
}
