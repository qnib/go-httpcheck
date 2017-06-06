package httpcheck

type Metric struct {
	Cnt float64
	ClientIP string
}


func NewMetric(ip string) Metric {
	return Metric{
		Cnt:      0.0,
		ClientIP: ip,
	}
}

func (m *Metric) Increment() {
	m.Cnt += 1.0
}