package openaigo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mockserver *httptest.Server

func TestMain(m *testing.M) {
	mockserver = testserverV1()
	code := m.Run()
	os.Exit(code)
}

func testserverV1() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/completions", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			f, e := os.Open("./testdata/completion.json")
			if e != nil {
				panic(e)
			}
			defer f.Close()
			io.Copy(w, f)
		}
	})
	mux.HandleFunc("/models", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			f, e := os.Open("./testdata/models_list.json")
			if e != nil {
				panic(e)
			}
			defer f.Close()
			io.Copy(w, f)
		}
	})
	mux.HandleFunc("/images/edits", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			json.NewEncoder(w).Encode(map[string]any{
				"created": 1670725494,
				"data": []any{
					map[string]any{"url": "https://otiai10.com/foobaa"},
				},
			})
		}
	})
	mux.HandleFunc("/images/variations", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			json.NewEncoder(w).Encode(map[string]any{
				"created": 1670725494,
				"data": []any{
					map[string]any{"url": "https://otiai10.com/foobaa"},
				},
			})
		}
	})
	mux.HandleFunc("/files", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(map[string]any{
				"object": "list",
				"data": []any{
					map[string]any{
						"id":         "file-ccdDZrC3iZVNiQVeEA6Z66wf",
						"object":     "file",
						"bytes":      175,
						"created_at": 1613677385,
						"filename":   "train.jsonl",
						"purpose":    "search",
					},
				},
			})
		case http.MethodPost:
			json.NewEncoder(w).Encode(map[string]any{
				"id":         "file-ccdDZrC3iZVNiQVeEA6Z66wf",
				"object":     "file",
				"bytes":      175,
				"created_at": 1613677385,
				"filename":   "train.jsonl",
				"purpose":    "search",
			})
		}
	})
	mux.HandleFunc("/files/abcdefg", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(map[string]any{
				"id":         "file-ccdDZrC3iZVNiQVeEA6Z66wf",
				"object":     "file",
				"bytes":      175,
				"created_at": 1613677385,
				"filename":   "train.jsonl",
				"purpose":    "search",
			})
		case http.MethodDelete:
			json.NewEncoder(w).Encode(map[string]any{
				"id":      "file-ccdDZrC3iZVNiQVeEA6Z66wf",
				"object":  "file",
				"deleted": true,
			})
		}
	})
	mux.HandleFunc("/files/abc", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]any{
				"error": map[string]any{
					"message": "No such File object: abc",
					"type":    "invalid_request_error",
					"param":   "id",
					"code":    nil,
				},
			})
		}
	})
	mux.HandleFunc("/files/zzz", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{....///"))
		}
	})
	mux.HandleFunc("/files/abcdefg/content", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			f, err := os.Open("./testdata/train.jsonl")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer f.Close()
			io.Copy(w, f)
		}
	})

	return httptest.NewServer(mux)
}
