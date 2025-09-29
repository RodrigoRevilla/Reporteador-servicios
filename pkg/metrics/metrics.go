package metrics

type Metrics interface {
	IncCounter(name string)
	ObserveHistogram(name string, value float64)
}

type NoOp struct{}

func (NoOp) IncCounter(name string)                      {}
func (NoOp) ObserveHistogram(name string, value float64) {}
