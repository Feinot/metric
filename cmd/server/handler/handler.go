package handler

import (
	"encoding/json"
	"github.com/Feinot/metric/cmd/server/Forms"
	"github.com/Feinot/metric/cmd/server/storage"
	"io/ioutil"
	"net/http"
	"time"
)

type Metric Forms.Metric

var m Metric

func HandleGuage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "", 400)

	} else {
		err = json.Unmarshal(body, &m)
		if err != nil {

			http.Error(w, "", http.StatusBadRequest)
		}
	}
	if m.MetricName == "" && m.MetricType != "guage" {
		http.Error(w, "gu", 404)

		return
	}
	storage.Storage.Guage[m.MetricName] = m.MetricValue

	http.Error(w, "gu", 200)

}
func HandleCaunter(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "", 400)
		return

	} else {
		err = json.Unmarshal(body, &m)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)

		}
	}

	if m.MetricName == "" && m.MetricType != "caunter" {
		http.Error(w, "", 400)

		return
	}
	s := make(map[time.Time][]map[string][]Forms.Monitor)
	q := make(map[string][]Forms.Monitor)
	q[m.MetricName] = append(q[m.MetricName], m.MetricValue)

	s[time.Now()] = append(storage.Storage.Counter[time.Now()], q)
	storage.Storage.Counter = s

	http.Error(w, "", 200)
}
