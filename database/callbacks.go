package database

import (
	"github.com/shixinshuiyou/framework/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	spanGormKey   = "opentracingSpan"
	operateCreate = "create"
	operateQuery  = "query"
	operateUpdate = "update"
	operateDelete = "delete"
)

type callbacks struct {
	serverName string
}

func (c *callbacks) before(db *gorm.DB) {
	_, sp := otel.Tracer(c.serverName).Start(db.Statement.Context, "db-"+db.Statement.Table)
	db.Set(spanGormKey, sp)
}

func (c *callbacks) after(db *gorm.DB, operate string) {
	val, ok := db.Get(spanGormKey)
	if !ok {
		return
	}
	sp := val.(trace.Span)

	sp.SetAttributes(
		attribute.String("db.table", db.Statement.Table),
		attribute.Int64("db.rowsAffected", db.Statement.DB.RowsAffected),
		attribute.String("sql", db.Statement.SQL.String()),
		attribute.String("operate", operate),
	)
	// sql执行失败
	if db.Error != nil {
		sp.SetAttributes()
	}
}

func (c *callbacks) registerCallbacks(db *gorm.DB, operate string) {
	beforeOperateName := "trace:before_" + operate
	afterOperateName := "trace:after_" + operate
	gormCallbackName := "gorm:" + operate
	log.Logger.Info("register")

	switch operate {
	case operateCreate:
		db.Callback().Create().Before(gormCallbackName).Register(beforeOperateName, func(db *gorm.DB) {
			c.before(db)
		})
		db.Callback().Create().After(gormCallbackName).Register(afterOperateName, func(db *gorm.DB) {
			c.after(db, operate)
		})
	case operateQuery:
		db.Callback().Query().Before(gormCallbackName).Register(beforeOperateName, func(db *gorm.DB) {
			log.Logger.Info("++++++++++++")
			c.before(db)
		})
		db.Callback().Query().After(gormCallbackName).Register(afterOperateName, func(db *gorm.DB) {
			c.after(db, operate)
			log.Logger.Info("-------------")
		})
	case operateUpdate:
		db.Callback().Update().Before(gormCallbackName).Register(beforeOperateName, func(db *gorm.DB) {
			c.before(db)
		})
		db.Callback().Update().After(gormCallbackName).Register(afterOperateName, func(db *gorm.DB) {
			c.after(db, operate)
		})
	case operateDelete:
		db.Callback().Delete().Before(gormCallbackName).Register(beforeOperateName, func(db *gorm.DB) {
			c.before(db)
		})
		db.Callback().Delete().After(gormCallbackName).Register(afterOperateName, func(db *gorm.DB) {
			c.after(db, operate)
		})
	}
}

// AddGormCallbacks 为当前db注入trace
func AddGormCallbacks(db *gorm.DB, name string) {
	c := &callbacks{serverName: name}
	c.registerCallbacks(db, operateCreate)
	c.registerCallbacks(db, operateQuery)
	c.registerCallbacks(db, operateUpdate)
	c.registerCallbacks(db, operateDelete)
}
