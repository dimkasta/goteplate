package goteplate

import (
	"bytes"
	"github.com/dimkasta/gologger"
	"html/template"
)

type TemplateService struct {
	logger     *gologger.LoggerService
	repository *SqliteTemplateRepository
}

func (service *TemplateService) Get(name string, data map[string]string) (string, error) {


	//html := "<html><body><h1>Hello World!<small>{{ .test }}</small></h1></body></html>"
	html := service.repository.Get(name)

	templ, err := template.New(name).Parse(string(html))

	if nil != err {
		service.logger.Error(err.Error())
	}

	var output bytes.Buffer

	err = templ.Execute(&output, data)

	if nil != err {
		service.logger.Error(err.Error())
	}

	return output.String(), err
}

func NewTemplateService(logger *gologger.LoggerService, repository *SqliteTemplateRepository) *TemplateService {
	return &TemplateService{
		logger: logger,
		repository: repository,
	}
}