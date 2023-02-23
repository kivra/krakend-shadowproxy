package shadowproxy

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/luraproject/lura/v2/config"
)

type jsonConfig struct {
	Host                     []string `json:"host"`
	URLPattern               string   `json:"url_pattern"`
	Method                   string   `json:"method"`
	HostSanitizationDisabled bool     `json:"disable_host_sanitize"`
	Timeout                  string   `json:"timeout"`
}

type ProxyConfig struct {
	Host                     []string
	URLPattern               string
	Method                   string
	HostSanitizationDisabled bool
	Timeout                  time.Duration
}

const Namespace = "kivra/shadowproxy"

func addNameSpace(s string) string {
	return Namespace + ": " + s
}

func configGetter(e *config.EndpointConfig) (*ProxyConfig, bool) {
	jsonCfg := new(jsonConfig)
	cfg := new(ProxyConfig)

	tmp, ok := e.ExtraConfig[Namespace]
	if !ok {
		return cfg, false
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(tmp); err != nil {
		panic(addNameSpace("Error: failed to parse proxy config: " + err.Error()))
	}
	if err := json.NewDecoder(buf).Decode(jsonCfg); err != nil {
		panic(addNameSpace("Error: failed to parse proxy config: " + err.Error()))
	}

	cfg.Host = jsonCfg.Host
	cfg.HostSanitizationDisabled = jsonCfg.HostSanitizationDisabled
	cfg.Method = jsonCfg.Method
	cfg.URLPattern = jsonCfg.URLPattern

	if jsonCfg.Timeout == "" {
		cfg.Timeout = e.Timeout
	} else {
		t, err := time.ParseDuration(jsonCfg.Timeout)
		if err != nil {
			panic(addNameSpace("Error: failed to parse proxy timeout config: " + err.Error()))
		}
		cfg.Timeout = t
	}

	if cfg.Method == "" {
		cfg.Method = "GET"
	}

	if !cfg.HostSanitizationDisabled {
		cfg.Host = config.NewURIParser().CleanHosts(cfg.Host)
	}

	return cfg, true
}
