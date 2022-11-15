package shadowproxy

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	urlPattern, urlKeys := parseURLPattern("/v1/{user}/foo/{bar}/{ukey}")
	expected := "/v1/{{.User}}/foo/{{.Bar}}/{{.Ukey}}"
	if urlPattern != expected {
		t.Fatalf("expected '%s', got '%s'", expected, urlPattern)
	}
	if urlKeys[0] != "User" { // nolint
		t.Fatalf("expected 'User', got '%s'", urlKeys[0])
	}
	if urlKeys[1] != "Bar" {
		t.Fatalf("expected 'Bar', got '%s'", urlKeys[1])
	}
	if urlKeys[2] != "Ukey" {
		t.Fatalf("expected 'Ukey', got '%s'", urlKeys[2])
	}
	if len(urlKeys) != 3 {
		t.Fatalf("expected len '3', got '%d'", len(urlKeys))
	}
}

func TestTemplateNoPathVars(t *testing.T) {
	urlPattern := "/v1/user/ukey"
	parsed, urlKeys := parseURLPattern(urlPattern)
	if parsed != urlPattern {
		t.Fatalf("expected '%s', got '%s'", urlPattern, parsed)
	}
	if len(urlKeys) != 0 {
		t.Fatalf("expected len '0', got '%d'", len(urlKeys))
	}
}

func TestTemplatePathVarsWithUpperCase(t *testing.T) {
	urlPattern, urlKeys := parseURLPattern("/v1/{User}/foo/{BAR}/{uKey}")
	expected := "/v1/{{.User}}/foo/{{.BAR}}/{{.UKey}}"
	if urlPattern != expected {
		t.Fatalf("expected '%s', got '%s'", expected, urlPattern)
	}
	if urlKeys[0] != "User" {
		t.Fatalf("expected 'User', got '%s'", urlKeys[0])
	}
	if urlKeys[1] != "BAR" {
		t.Fatalf("expected 'BAR', got '%s'", urlKeys[1])
	}
	if urlKeys[2] != "UKey" {
		t.Fatalf("expected 'UKey', got '%s'", urlKeys[2])
	}
	if len(urlKeys) != 3 {
		t.Fatalf("expected len '3', got '%d'", len(urlKeys))
	}
}
