package db

import (
	"fmt"
	"gorm.io/gorm"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const spanGormKey = "opentracingSpan"

type callbacks struct{}

func newCallbacks() *callbacks {
	return &callbacks{}
}

func (c *callbacks) beforeCreate(db *gorm.DB) { c.before(db) }
func (c *callbacks) afterCreate(db *gorm.DB)  { c.after(db, "INSERT") }
func (c *callbacks) beforeQuery(db *gorm.DB)  { c.before(db) }
func (c *callbacks) afterQuery(db *gorm.DB)   { c.after(db, "SELECT") }
func (c *callbacks) beforeUpdate(db *gorm.DB) { c.before(db) }
func (c *callbacks) afterUpdate(db *gorm.DB)  { c.after(db, "UPDATE") }
func (c *callbacks) beforeDelete(db *gorm.DB) { c.before(db) }
func (c *callbacks) afterDelete(db *gorm.DB)  { c.after(db, "DELETE") }

func (c *callbacks) before(db *gorm.DB) {
	_, sp := otel.Tracer("appconfig.ServerSetting.ServerName").Start(db.Statement.Context, "db")
	db.Set(spanGormKey, sp)
}

func (c *callbacks) after(db *gorm.DB, operation string) {
	val, ok := db.Get(spanGormKey)
	if !ok {
		return
	}
	sp := val.(trace.Span)
	if db.Error != nil {
		sp.SetStatus(codes.Error, db.Error.Error())
	}
	sp.SetAttributes(attribute.String("db.table", db.Statement.Table))
	sp.SetAttributes(attribute.Int64("db.rowsAffected", db.Statement.DB.RowsAffected))
	sp.SetAttributes(attribute.String("sql", db.Statement.SQL.String()))
	sp.SetAttributes(attribute.String("operation", operation))
	sp.End()
}

func AddGormCallbacks(db *gorm.DB) {
	cbs := newCallbacks()
	registerCallbacks(db, "create", cbs)
	registerCallbacks(db, "query", cbs)
	registerCallbacks(db, "update", cbs)
	registerCallbacks(db, "delete", cbs)
	registerCallbacks(db, "row_query", cbs)
}

func registerCallbacks(db *gorm.DB, name string, c *callbacks) {
	beforeName := fmt.Sprintf("tracing:%v_before", name)
	afterName := fmt.Sprintf("tracing:%v_after", name)
	gormCallbackName := fmt.Sprintf("gorm:%v", name)
	// gorm does some magic, if you pass CallbackProcessor here - nothing works
	switch name {
	case "create":
		db.Callback().Create().Before(gormCallbackName).Register(beforeName, c.beforeCreate)
		db.Callback().Create().After(gormCallbackName).Register(afterName, c.afterCreate)
	case "query":
		db.Callback().Query().Before(gormCallbackName).Register(beforeName, c.beforeQuery)
		db.Callback().Query().After(gormCallbackName).Register(afterName, c.afterQuery)
	case "update":
		db.Callback().Update().Before(gormCallbackName).Register(beforeName, c.beforeUpdate)
		db.Callback().Update().After(gormCallbackName).Register(afterName, c.afterUpdate)
	case "delete":
		db.Callback().Delete().Before(gormCallbackName).Register(beforeName, c.beforeDelete)
		db.Callback().Delete().After(gormCallbackName).Register(afterName, c.afterDelete)
	}
}
