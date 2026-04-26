package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDriver struct{}
type sqliteDB struct {
	db *gorm.DB
}

func init() {
	Register("sqlite", &sqliteDriver{})
}

func (d *sqliteDriver) Open(config map[string]string) (DB, error) {
	path := config["path"]
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &sqliteDB{db: db}, nil
}

func (s *sqliteDB) DB() *gorm.DB {
	return s.db
}

func (s *sqliteDB) Close() error {
	sqlDB, _ := s.db.DB()
	return sqlDB.Close()
}
