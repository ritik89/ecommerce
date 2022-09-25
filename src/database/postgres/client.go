package postgres

import (
	"context"
	"fmt"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpgx"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	singleton  sync.Once
	postgresDB = &gorm.DB{}
)

func InitDB() *gorm.DB {
	var err error
	singleton.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			viper.GetString("database.postgres.source.host"),
			viper.GetString("database.postgres.source.user"),
			viper.GetString("database.postgres.source.password"),
			viper.GetString("database.postgres.source.db_name"),
			viper.GetInt("database.postgres.source.port"),
		)
		postgresDB, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "nrpgx",
			DSN:        dsn,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(fmt.Sprintf("not able to connect to the database. Error:- %s", err.Error()))
		}
	})
	return postgresDB
}

func GetDB(rCtx context.Context) *gorm.DB {
	if rCtx != nil {
		return postgresDB.WithContext(newrelic.NewContext(context.Background(), newrelic.FromContext(rCtx)))
	}
	return postgresDB
}
