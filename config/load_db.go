package config

import (
	"database/sql"
	"errors"
	"go-boiler-plate/db"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitializeDB(config *AppConfig) db.Store {
	dbSource := viper.GetString("DB_SOURCE")
	conn, err := sql.Open(config.DBDriver, dbSource)
	if err != nil {
		log.Fatal("Failed to load db", err)
	}
	conn.SetMaxIdleConns(config.DBMaxIdleConn)
	conn.SetMaxOpenConns(config.DBMaxOpenConn)
	conn.SetConnMaxLifetime(time.Hour)
	migrationURL := rootPath() + "/db/migration"
	runDBMigration("file://"+migrationURL, dbSource)
	return db.NewStore(conn)
}

func runDBMigration(migrateURL string, dbSource string) {
	migration, err := migrate.New(migrateURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to run migrate up:", err)
	}
}
