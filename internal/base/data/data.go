package data

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	ormlog "xorm.io/xorm/log"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(db *xorm.Engine) (*Data, func(), error) {
	cleanup := func() {
		db.Close()
	}
	return &Data{
		DB: db,
	}, cleanup, nil
}

func NewDB(debug bool, dataConf *Database) (*xorm.Engine, error) {
	dataConf.Driver = "mysql"
	engine, err := xorm.NewEngine(dataConf.Driver, dataConf.Connection)
	if err != nil {
		return nil, err
	}

	if debug {
		engine.ShowSQL(true)
	} else {
		engine.SetLogLevel(ormlog.LOG_ERR)
	}

	if err = engine.Ping(); err != nil {
		return nil, err
	}

	if dataConf.ConnMaxLifeTime > 0 {
		engine.SetConnMaxLifetime(time.Duration(dataConf.ConnMaxLifeTime) * time.Second)
	}
	if dataConf.MaxOpenConn > 0 {
		engine.SetMaxOpenConns(dataConf.MaxOpenConn)
	}
	if dataConf.MaxIdleConn > 0 {
		engine.SetMaxIdleConns(dataConf.MaxIdleConn)
	}

	return engine, nil
}
