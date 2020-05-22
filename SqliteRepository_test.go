package goteplate

import (
	"github.com/dimkasta/gologger"
	"testing"
)

func TestTemplateService_Get(t *testing.T) {
	logger := gologger.NewLoggerService()
	service := NewSqliteTemplateRepository("templates.db", logger)

	html := service.Get("test")

	if "<h1>test: {{ .test }} </h1>" != html {
		t.Errorf("Template should be test. existing: %s", html)
	}
}
