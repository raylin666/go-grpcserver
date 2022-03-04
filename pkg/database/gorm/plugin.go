package gorm

import (
	"gorm.io/gorm"
	"time"
)

const (
	callBackBeforeName = "database:before"
	callBackAfterName  = "database:after"
	startTime          = "_start_time"
)

type TracePlugin struct {
	Before func(gormDb *gorm.DB)
	After  func(gormDb *gorm.DB, sql string, ts time.Time)
}

func (op *TracePlugin) Name() string {
	return "tracePlugin"
}

func (op *TracePlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	_ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, op.before)
	_ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, op.before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, op.before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, op.before)
	_ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, op.before)
	_ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, op.before)

	// 结束后
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, op.after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, op.after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, op.after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, op.after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, op.after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, op.after)
	return
}

var _ gorm.Plugin = &TracePlugin{}

func (op *TracePlugin) before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	if op.Before != nil {
		op.Before(db)
	}
	return
}

func (op *TracePlugin) after(db *gorm.DB) {
	_ts, isExist := db.InstanceGet(startTime)
	if !isExist {
		return
	}

	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}

	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	if op.After != nil {
		op.After(db, sql, ts)
	}

	return
}
