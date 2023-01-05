package tool

var serviceTemplate = `
package service

import (
	"showapi/repository"
)

type  I{{class}}Service interface {
}

type {{class}}Service struct {
	*BaseService
	{{obj}}Repository    repository.I{{class}}Repository
}

func New{{class}}Service(option ...Option) *{{class}}Service {
	return &{{class}}Service{BaseService:NewByOption(option...)}
}

func  ({{obj}}Service *{{class}}Service)Set{{class}}Repository(repository repository.I{{class}}Repository){
	{{obj}}Service.{{obj}}Repository = repository
}

func ({{obj}}Service *{{class}}Service) Init() {
	options := []repository.Option{repository.WithContext({{obj}}Service.ctx), repository.WithLog({{obj}}Service.log)}
	{{obj}}Repository := repository.New{{class}}Repository(options...)
	{{obj}}Service.Set{{class}}Repository({{obj}}Repository)
}

func ({{obj}}Service *{{class}}Service) Detail() {

}

func ({{obj}}Service *{{class}}Service) Delete(id int) bool {
	return false
}

func ({{obj}}Service *{{class}}Service) List() {

}

func ({{obj}}Service *{{class}}Service) Update(id int) {

}


`
