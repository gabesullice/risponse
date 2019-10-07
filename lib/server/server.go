package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func ListenAndServe(addr string, config configuration) {
	mux := http.NewServeMux()
	for _, resource := range config.Resources {
		mux.HandleFunc(resource.Path, getResourceHandlerFunc(resource, config))
	}
	http.ListenAndServe(addr, mux)
}

func getResourceHandlerFunc(resource resource, config configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			addHeaders(w, resource.allowHeaders(), resource.Cors.headers(r), config.Defaults.Cors.headers(r))
			w.WriteHeader(http.StatusNoContent)
			return
		}
		for _, method := range resource.Methods {
			if method == r.Method {
				addHeaders(w, resource.Cors.headers(r), config.Defaults.Cors.headers(r), resource.Headers, config.Defaults.Headers)
				w.WriteHeader(resource.Status)
				filename := fmt.Sprintf(".%s/%s.json", resource.Path, strings.ToLower(method))
				f, err := os.Open(filename)
				defer f.Close()
				if err == nil {
					io.Copy(w, f)
				}
				return
			}
		}
		addHeaders(w, resource.allowHeaders())
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func addHeaders(w http.ResponseWriter, addl ...map[string]string) {
	headers := w.Header()
	for _, add := range addl {
		for name, value := range add {
			if _, ok := headers[name]; !ok {
				headers.Add(name, value)
			}
		}
	}
}
