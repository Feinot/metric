package handler

import (
	"fmt"
	"github.com/Feinot/metric/cmd/server/form"
	"net/http"
	"strconv"
	"strings"
)

var storage form.MemStorage

type Metric form.Metric

var m Metric

func HandleGuage(w http.ResponseWriter) {
	if m.MetricName == "" {
		http.Error(w, "", 400)
		return
	}
	storage.Guage[m.MetricName] = m.Guage
	fmt.Println("type = ", m.MetricType, " name = ", m.MetricName, " Value = ", m.Guage)
	http.Error(w, "", 200)

}
func HandleCaunter(w http.ResponseWriter) {
	if m.MetricName == "" {
		http.Error(w, "", 400)
		return
	}
	s := make(map[string][]int64)
	s[m.MetricName] = append(storage.Counter[m.MetricName], m.Counter)
	storage.Counter = s
	fmt.Println("type = ", m.MetricType, " name = ", m.MetricName, " Value = ", m.Counter)
	http.Error(w, "", 200)
}

func RequestHandle(w http.ResponseWriter, r *http.Request) {
	var err error
	arr := make([]string, 3, 3)
	url := strings.Split(r.URL.Path, "/update/")
	url = strings.Split(url[1], "/")
	fmt.Println(url)
	for q := 0; q < len(url); q++ {
		arr[q] = url[q]
	}
	m.MetricType = arr[0]
	m.MetricName = arr[1]

	switch m.MetricType {
	case "gauge":
		m.Guage, err = strconv.ParseFloat(arr[2], 64)
		if err != nil {
			http.Error(w, "", 400)
			return
		}
		HandleGuage(w)
	case "counter":

		m.Counter, err = strconv.ParseInt(arr[2], 10, 64)
		if err != nil {
			http.Error(w, "", 400)
			return
		}
		HandleCaunter(w)
	default:

		http.Error(w, "", 400)
	}
}
