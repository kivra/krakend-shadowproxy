package shadowproxy

import (
	"bytes"
	"encoding/json"

	"github.com/luraproject/lura/v2/config"
)

type ProxyConfig struct {
	Host                     []string `json:"host"`
	URLPattern               string   `json:"url_pattern"`
	Method                   string   `json:"method"`
	HostSanitizationDisabled bool     `json:"disable_host_sanitize"`
}

const Namespace = "kivra/shadowproxy"

func addNameSpace(s string) string {
	return Namespace + ": " + s
}

func configGetter(e config.ExtraConfig) (*ProxyConfig, bool) {
	cfg := new(ProxyConfig)

	tmp, ok := e[Namespace]
	if !ok {
		return cfg, false
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(tmp); err != nil {
		panic(addNameSpace("Error: failed to parse proxy config: " + err.Error()))
	}
	if err := json.NewDecoder(buf).Decode(cfg); err != nil {
		panic(addNameSpace("Error: failed to parse proxy config: " + err.Error()))
	}

	if cfg.Method == "" {
		cfg.Method = "GET"
	}

	if !cfg.HostSanitizationDisabled {
		cfg.Host = config.NewURIParser().CleanHosts(cfg.Host)
	}

	return cfg, true
}
