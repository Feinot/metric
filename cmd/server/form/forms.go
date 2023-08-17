package form

type Metric struct {
	MetricType  string `json:"MetricType"`
	MetricName  string `json:"MetricName"`
	MetricValue uint64 `json:"MetricValue"`
}
type MemStorage struct {
	Guage   interface{}
	Counter map[string][]interface{}
}
