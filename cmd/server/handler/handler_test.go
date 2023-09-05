package handler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {

		return nil
	},
}

/*
func TestMetric_HandleCaunter(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
		url         string
		requestType string
		metricValue string
		metricType  string
		metricName  string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "text/plain",
				url:         "/update/",
				requestType: "POST",
				metricType:  "gauge/",
				metricName:  "GOOS/",
				metricValue: "123",
			},
		},

		{
			name: "positive test #2",
			want: want{
				code: 200,
				//response:    `{"status":"ok"}`,
				contentType: "text/plain",
				url:         "/update/",
				requestType: "POST",
				metricType:  "counter/",
				metricName:  "GOOS/",
				metricValue: "123",
			},
		},

		{
			name: "negative test #1",
			want: want{
				code: 400,

				contentType: "application/json",
				url:         "/update/",
				requestType: "POST",
				metricType:  "uncown/",
				metricName:  "GOOS/",
				metricValue: "123",
			},
		},
		{
			name: "negative test #2",
			want: want{
				code: 404,

				contentType: "application/json",
				url:         "/update/",
				requestType: "POST",
				metricType:  "counter/",
				metricName:  "/",
				metricValue: "123",
			},
		},
		{
			name: "negative test #2",
			want: want{
				code: 400,

				contentType: "application/json",
				url:         "/update/",
				requestType: "POST",
				metricType:  "counter/",
				metricName:  "GOOS/",
				metricValue: " ",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//test.want.url = fmt.Sprintf("%s%s%s%s", test.want.url, test.want.metricType, test.want.metricName, test.want.metricValue)

			res, _ := client.Post(fmt.Sprintf("%s%s%s%s%s", "http://localhost:8080", test.want.url, test.want.metricType, test.want.metricName, test.want.metricValue), "text/plain", nil)
			res.Body.Close()

			assert.Equal(t, res.StatusCode, test.want.code)

		})
	}

}

*/

type routeTest struct {
	title string // title of the test
	//route           *Route            // the route being tested
	types           string            // a request to test the route
	vars            map[string]string // the expected vars of the match
	scheme          string            // the expected scheme of the built URL
	host            string            // the expected host of the built URL
	path            string            // the expected path of the built URL
	query           string            // the expected query string of the built URL
	pathTemplate    string            // the expected path template of the route
	hostTemplate    string            // the expected host template of the route
	queriesTemplate string            // the expected query template of the route
	methods         []string          // the expected route methods
	pathRegexp      string            // the expected path regexp
	queriesRegexp   string            // the expected query regexp
	shouldMatch     bool              // whether the request is expected to match the route at all
	shouldRedirect  bool              // whether the request should result in a redirect
	statusCode      int
}

func TestMetric_HandleCaunter(t *testing.T) {
	tests := []routeTest{
		{
			title:      "Positive test#1 guage ",
			types:      "POST",
			vars:       map[string]string{"type": "guage", "name": "asd", "value": "123"},
			host:       "http://localhost:8080/update/gauge/asd/123",
			statusCode: 200,
		},
		{
			title:      "Positive test#2 ",
			types:      "POST",
			vars:       map[string]string{"type": "counter", "name": "asd", "value": "123"},
			host:       "http://localhost:8080/update/gauge/asd/123",
			statusCode: 200,
		},
		{
			title:      "Nigative test#1 ",
			types:      "POST",
			vars:       map[string]string{"type": "counter", "name": "asd", "value": ""},
			host:       "http://localhost:8080/update/gauge/asd/",
			statusCode: 400,
		},
		{
			title:      "Nigative test#2 ",
			types:      "POST",
			vars:       map[string]string{"type": "counter", "name": "", "value": "12"},
			host:       "http://localhost:8080/update/gauge/asd/",
			statusCode: 400,
		},
		{
			title:      "Nigative test#3 ",
			types:      "POST",
			vars:       map[string]string{"type": "counter", "name": "", "value": "12"},
			host:       "http://localhost:8080/update/gauge//132",
			statusCode: 404,
		},
		{
			title:      "Nigative test#4 ",
			types:      "POST",
			vars:       map[string]string{"type": "counter", "name": "", "value": "12"},
			host:       "http://localhost:8080/update//asd/132",
			statusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {

			r, err := http.NewRequest(test.types, test.host, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			w := httptest.NewRecorder()

			//Hack to try to fake gorilla/mux vars

			// CHANGE THIS LINE!!!

			RequestUpdateHandle(w, r)

			assert.Equal(t, test.statusCode, w.Code)
			//assert.Equal(t, []byte("abcd"), w.Body.Bytes())
		})
	}
}
