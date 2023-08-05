package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMetric_HandleCaunter(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
		url         string
		requestType string
		metric      *strings.Reader
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
				contentType: "application/json",
				url:         "/update/counter",
				requestType: "POST",
				metric:      strings.NewReader(`{ "MetricType": "counter", "MetricName": "GOOS", "MetricValue": 123}`),
			},
		},
		{
			name: "positive test #2",
			want: want{
				code:        200,
				response:    `{"status":"ok"}`,
				contentType: "application/json",
				url:         "/update/guage",
				requestType: "POST",
				metric:      strings.NewReader(`{ "MetricType": "counter", "MetricName": "GOOS", "MetricValue": 123}`),
			},
		},
		{
			name: "negative test #1",
			want: want{
				code: 400,

				contentType: "application/json",
				url:         "/update/counter",
				requestType: "POST",
				metric:      strings.NewReader(`{ "MetricType": "asd", "MetricName": "GOOS", "MetricValue": 123}`),
			},
		},
		{
			name: "negative test #2",
			want: want{
				code: 404,

				contentType: "application/json",
				url:         "/update/counter",
				requestType: "POST",
				metric:      strings.NewReader(`{ "MetricType": "counter", "MetricName": "", "MetricValue": 123}`),
			},
		},
		{
			name: "negative test #2",
			want: want{
				code: 404,

				contentType: "application/json",
				url:         "/update/counter",
				requestType: "POST",
				metric:      strings.NewReader(`{ "MetricType": "counter", "MetricName": "", "MetricValue": }`),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			request := httptest.NewRequest(test.want.requestType, test.want.url, test.want.metric)

			w := httptest.NewRecorder()
			RequestHandle(w, request)

			res := w.Result()

			assert.Equal(t, res.StatusCode, test.want.code)

		})
	}

}
