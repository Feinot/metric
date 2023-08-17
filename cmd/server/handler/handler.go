package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Feinot/metric/cmd/server/forms"
	"io"
	"net/http"
)

var storage forms.MemStorage

type Metric forms.Metric

var m Metric

func HandleGuage(w http.ResponseWriter) {
	if m.MetricName == "" {
		http.Error(w, "", 404)
		return
	}
	storage.Guage = m.MetricValue
	http.Error(w, "", 200)

}
func HandleCaunter(w http.ResponseWriter) {
	if m.MetricName == "" {
		http.Error(w, "", 404)
		return
	}
	s := make(map[string][]interface{})
	s[m.MetricName] = append(storage.Counter[m.MetricName], m.MetricValue)
	storage.Counter = s
	http.Error(w, "", 200)
}

func RequestHandle(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "", 400)
		//fmt.Fprintf(w, "err %q\n", err, err.Error())

	} else {
		err = json.Unmarshal(body, &m)
		if err != nil {
			fmt.Println(err.Error())

		}
	}
	switch m.MetricType {
	case "gauge":
		HandleGuage(w)
	case "counter":
		HandleCaunter(w)
	default:
		http.Error(w, "", 400)
	}
}
