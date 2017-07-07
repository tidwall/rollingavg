package rollingavg

type RollingAvg struct {
	values []float64
	index  int
	count  int
	total  float64
}

func NewWindowSize(window int) *RollingAvg {
	return &RollingAvg{
		values: make([]float64, window),
	}
}
func New() *RollingAvg {
	return NewWindowSize(256)
}
func (a *RollingAvg) Add(value float64) float64 {
	prev := a.values[a.index]
	a.values[a.index] = value
	a.total = a.total - prev + value
	a.index++
	if a.index == len(a.values) {
		a.index = 0
	}
	if a.count < len(a.values) {
		a.count++
	}
	return float64(a.total) / float64(a.count)
}
