package database

import (
	"context"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	instance *pg.DB
	once     sync.Once
)

func NewPgDB(pgOptions *pg.Options) *pg.DB {
	once.Do(func() {
		instance = Connect(pgOptions)
	})

	return instance
}

func Connect(pgOptions *pg.Options) *pg.DB {
	log.Info("connecting to postgres database...")
	ctx := context.Background()
	db := pg.Connect(pgOptions)

	log.Info("verifying postgres connection...")
	if err := db.Ping(ctx); err != nil {
		log.Errorln(err)
		time.Sleep(2 * time.Second)

		db = pg.Connect(pgOptions)
		if err := db.Ping(ctx); err == nil {
			log.Info("successfully connected to postgres database...")
			return db
		}

		log.Panicf("Postgres connection error %+v\n", err)
	}

	log.Info("Connection to postgres verified...")
	return db
}
