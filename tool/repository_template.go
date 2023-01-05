package tool

var repositoryTemplate = `
package repository

import (
	"showapi/models"
	"gorm.io/gorm/clause"
)

type  I{{class}}Repository interface {
	GetKpi(id int)({{obj}} *models.{{class}}, err error)
	Add({{obj}} *models.{{class}})(err error)
	Update(old *models.{{class}}, {{obj}} *models.{{class}})(err error) 
	UpdateWithParams({{obj}} *models.{{class}},params map[string]interface{})(err error)
}



type {{class}}Repository struct {
	*BaseRepository
}

func New{{class}}Repository(opts ...Option) *{{class}}Repository {
	baseRepository := NewByOption(opts...)
	{{obj}}  := {{class}}Repository{baseRepository}
	return &{{obj}}
}
func ({{obj}}Repository *{{class}}Repository)SetBaseRepository(repository *BaseRepository)  {
	{{obj}}Repository.BaseRepository =  repository
}

func ({{obj}}Repository *{{class}}Repository)GetKpi(id int)({{obj}} *models.{{class}}, err error)  {
	err = {{obj}}Repository.db.Where(" id = ?",  id).First(&{{obj}}).Error
	return
}
   
func ({{obj}}Repository *{{class}}Repository)Add({{obj}} *models.{{class}})(err error)  {
	err = {{obj}}Repository.db.Omit(clause.Associations).Create(&{{obj}}).Error
	return
}

func ({{obj}}Repository *{{class}}Repository)Update(old *models.{{class}}, {{obj}} *models.{{class}})(err error)  {
	err = {{obj}}Repository.db.Model(&old).Omit(clause.Associations).Save(&{{obj}}).Error
	return
}
     
func ({{obj}}Repository *{{class}}Repository)UpdateWithParams({{obj}} *models.{{class}},params map[string]interface{})(err error)  {
	err = {{obj}}Repository.db.Model(&{{obj}}).Omit(clause.Associations).Updates(params).Error
	return
}


`
