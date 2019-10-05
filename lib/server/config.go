package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type resource struct {
	Path    string            `json:"path"`
	Methods []string          `json:"methods"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Cors    cors              `json:"cors"`
}

func (r resource) allowHeaders() map[string]string {
	return map[string]string{
		"allow": strings.Join(r.Methods, ", "),
	}
}

type cors struct {
	AllowOrigin      []string `json:"allowOrigin"`
	AllowCredentials bool     `json:"allowCredentials"`
	ExposeHeaders    []string `json:"exposeHeaders"`
}

func (c cors) headers(r *http.Request) map[string]string {
	headers := map[string]string{}
	if len(c.AllowOrigin) >= 1 {
		for _, origin := range c.AllowOrigin {
			if origin == "*" || origin == r.Header.Get("origin") {
				headers["access-control-allow-origin"] = origin
			}
		}
	}
	if c.AllowCredentials {
		headers["access-control-allow-credentials"] = "true"
	}
	if len(c.ExposeHeaders) >= 1 {
		headers["access-control-expose-headers"] = strings.Join(c.ExposeHeaders, ", ")
	}
	return headers
}

type configuration struct {
	Defaults struct {
		Cors cors `json:"cors"`
	} `json:"defaults"`
	Resources []resource `json:"resources"`
}

func defaultConfig() configuration {
	var def configuration
	err := json.Unmarshal([]byte(`{}`), &def)
	if err != nil {
		panic(err)
	}
	return def
}

func LoadConfigFromFile(filename string) configuration {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	config := defaultConfig()
	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	for k, v := range config.Resources {
		if len(v.Methods) == 0 {
			config.Resources[k].Methods = append(v.Methods, "HEAD", "OPTIONS", "GET")
		}
	}
	return config
}
