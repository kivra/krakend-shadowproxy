package shadowproxy

import (
	"regexp"
	"strings"

	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/encoding"
	"github.com/luraproject/lura/v2/proxy"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type shadowFactory struct {
	f proxy.Factory
}

func (s shadowFactory) New(cfg *config.EndpointConfig) (p proxy.Proxy, err error) {
	if len(cfg.Backend) == 0 {
		err = proxy.ErrNoBackends
		return
	}
	p, err = s.f.New(cfg)
	if prxCfg, ok := configGetter(cfg.ExtraConfig); ok {
		pShadow, _ := s.f.New(shadowConfig(*cfg, *prxCfg))
		p = proxy.ShadowMiddleware(p, pShadow)
	}
	return
}

func NewShadowFactory(f proxy.Factory) proxy.Factory {
	return shadowFactory{f}
}

func shadowConfig(cfg config.EndpointConfig, prxCfg ProxyConfig) *config.EndpointConfig { // nolint
	urlPattern, urlKeys := parseURLPattern(prxCfg.URLPattern)
	cfg.Backend = []*config.Backend{{
		Host:                     prxCfg.Host,
		Method:                   prxCfg.Method,
		URLPattern:               urlPattern,
		HostSanitizationDisabled: prxCfg.HostSanitizationDisabled,
		URLKeys:                  urlKeys,
		Encoding:                 encoding.NOOP,
		Timeout:                  cfg.Timeout,
	}}
	return &cfg
}

func parseURLPattern(urlPattern string) (string, []string) { // nolint
	re := regexp.MustCompile(`/\{([a-zA-Z\-_0-9]+)\}`)
	urlKeys := []string{}
	if matches := re.FindAllStringSubmatch(urlPattern, -1); matches != nil {
		title := cases.Title(language.Und)
		for _, v := range matches {
			key := v[1]
			titleKey := title.String(key[:1]) + key[1:]
			urlPattern = strings.ReplaceAll(urlPattern, "{"+key+"}", "{{."+titleKey+"}}")
			urlKeys = append(urlKeys, titleKey)
		}
	}
	return urlPattern, urlKeys
}
