package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type baseHandler struct{}

func (h baseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	record := requestRecord{}
	record.TargetURL = r.URL.String()
	record.TargetHost = r.Host
	record.Headers = make(map[string]string, len(r.Header))
	for header, value := range r.Header {
		record.Headers[header] = value[0]
	}

	recordJSON, _ := json.Marshal(&record)

	buf := new(bytes.Buffer)
	json.Indent(buf, recordJSON, "", "\t")
	buf.WriteTo(w)
}

type requestRecord struct {
	TargetURL  string            `json:"target_url"`
	TargetHost string            `json:"target_host"`
	Headers    map[string]string `json:"request_headers"`
}

func main() {
	var handler http.Handler
	handler = baseHandler{}
	http.ListenAndServe(":8000", handler)
}
