package handler

import (
	"fmt"
	"github.com/Feinot/metric/forms"
	"github.com/Feinot/metric/storage"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Metric forms.Metric

var m Metric

func HandleGuage(w http.ResponseWriter) {

	s := make(map[string]float64)
	s[m.MetricName] = m.Guage
	storage.Storage.Guage = s
	http.Error(w, "", 200)

}
func HandleCaunter(w http.ResponseWriter) {

	s := make(map[string][]int64)
	s[m.MetricName] = append(storage.Storage.Counter[m.MetricName], m.Counter)
	storage.Storage.Counter = s
	http.Error(w, "", 200)
}

func RequestUpdateHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		var err error
		arr := mux.Vars(r)
		m.MetricType = arr["type"]
		m.MetricName = arr["name"]
		if m.MetricName == "" {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		log.Println(m.MetricType, "  ", m.MetricName)
		switch m.MetricType {
		case "gauge":
			m.Guage, err = strconv.ParseFloat(arr["value"], 64)
			if err != nil {
				http.Error(w, "", http.StatusBadRequest)
				return
			}
			HandleGuage(w)
		case "counter":

			m.Counter, err = strconv.ParseInt(arr["value"], 10, 64)
			if err != nil {
				http.Error(w, "", http.StatusBadRequest)
				return
			}
			HandleCaunter(w)
		default:

			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}
}
func RequestValueHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		arr := mux.Vars(r)
		m.MetricType = arr["type"]
		m.MetricName = arr["name"]
		if m.MetricName == "" {
			http.Error(w, arr["type"], http.StatusNotFound)
			return
		}
		switch m.MetricType {
		case "gauge":
			q := storage.Storage.Guage[m.MetricName]
			if q == 0 {
				http.Error(w, "", http.StatusNotFound)
				return
			}
			fmt.Fprintf(w, strconv.FormatFloat(q, 'f', 6, 64))
			http.Error(w, "", http.StatusOK)
		case "counter":
			q := storage.Storage.Counter[m.MetricName]
			if len(q) == 0 {
				http.Error(w, "", http.StatusNotFound)
				return

			}
			str := strconv.FormatInt(q[0], 10)
			for i := 1; i < len(q); i++ {

				str += "," + strconv.FormatInt(q[i], 10)
			}

			fmt.Fprintf(w, str)
			http.Error(w, "", http.StatusOK)
		default:
			http.Error(w, "", http.StatusNotFound)
			return
		}
	case http.MethodPost:

		return

	}
}
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl := template.Must(template.New("storage.Storage").Parse(`<div>
            <h1>Guage</h1>
			<p1>{{ .Guage}}</p>
			<h1>Counter</h1>
            <p>{{ .Counter}}</p>
        </div>`))
		tmpl.Execute(w, storage.Storage)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
