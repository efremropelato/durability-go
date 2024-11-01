package factorial

import (
	"math/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

type Factor struct {
	Mean              []float64
	Max               []float64
	Min               []float64
	Median            []float64
	Variance          []float64
	StandardDeviation []float64
	Skewness          []float64
	Type              int
	TypeName          string
	Values            [][]float64
}

func (f *Factor) SetFactor(Type int, arg1, arg2, arg3 float64, N, M int) {
	f.Type = Type
	f.setValues(arg1, arg2, arg3, N, M)
}

func (f *Factor) setValues(arg1, arg2, arg3 float64, N, M int) {
	f.Mean = make([]float64, M)
	f.Max = make([]float64, M)
	f.Min = make([]float64, M)
	f.Median = make([]float64, M)
	f.Variance = make([]float64, M)
	f.StandardDeviation = make([]float64, M)
	f.Skewness = make([]float64, M)

	for i := 0; i < M; i++ {
		switch f.Type {
		case 1:
			f.TypeName = "Beta Distribution"
			dist := distuv.Beta{
				Alpha: arg1,
				Beta:  arg2,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 2:
			f.TypeName = "Bernoulli Distribution"
			dist := distuv.Bernoulli{
				P: arg1,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 3:
			f.TypeName = "ChiSquare Distribution"
			dist := distuv.ChiSquared{
				K: arg1,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 4:
			f.TypeName = "Exponential Distribution"
			dist := distuv.Exponential{
				Rate: arg1,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 5:
			f.TypeName = "F Distribution"
			dist := distuv.F{
				D1: arg1,
				D2: arg2,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 6:
			f.TypeName = "Gamma Distribution"
			dist := distuv.Gamma{
				Alpha: arg1,
				Beta:  arg2,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 7:
			f.TypeName = "Normal Distribution"
			dist := distuv.Normal{
				Mu:    arg1,
				Sigma: arg2,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 8:
			f.TypeName = "Poisson Distribution"
			dist := distuv.Poisson{
				Lambda: arg1,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 9:
			f.TypeName = "T Distribution"
			dist := distuv.StudentsT{
				Nu: arg1,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 10:
			f.TypeName = "Uniform Distribution"
			for j := 0; j < N; j++ {
				f.Values[i][j] = rand.Float64()*(arg2-arg1) + arg1
			}
		case 11:
			f.TypeName = "Weibull Distribution"
			dist := distuv.Weibull{
				Lambda: arg1,
				K:      arg2,
			}
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		case 12:
			f.TypeName = "Triangular Distribution"
			dist := distuv.NewTriangle(arg1, arg2, arg3, nil)
			for j := 0; j < N; j++ {
				f.Values[i][j] = dist.Rand()
			}
		}
		f.Mean[i] = Mean(f.Values[i])
		f.Max[i] = Max(f.Values[i])
		f.Min[i] = Min(f.Values[i])
		f.Median[i] = Median(f.Values[i])
		f.Variance[i] = Variance(f.Values[i])
		f.StandardDeviation[i] = StdDev(f.Values[i])
		f.Skewness[i] = Skewness(f.Values[i])
	}
}

// Getter methods
func (f *Factor) GetAllValues() [][]float64       { return f.Values }
func (f *Factor) GetValues(i int) []float64       { return f.Values[i] }
func (f *Factor) GetMean() []float64              { return f.Mean }
func (f *Factor) GetMax() []float64               { return f.Max }
func (f *Factor) GetMin() []float64               { return f.Min }
func (f *Factor) GetMedian() []float64            { return f.Median }
func (f *Factor) GetVariance() []float64          { return f.Variance }
func (f *Factor) GetStandardDeviation() []float64 { return f.StandardDeviation }
func (f *Factor) GetSkewness() []float64          { return f.Skewness }
func (f *Factor) GetType() int                    { return f.Type }
func (f *Factor) GetTypeName() string             { return f.TypeName }
