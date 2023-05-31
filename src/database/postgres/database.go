package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"yandex-team.ru/bstask/database/model"
)

type Params struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func New(p *Params) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		p.Host, p.Port, p.User, p.Password, p.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	migrator := db.Migrator()
	tables := []any{model.Order{}, model.Courier{}}

	for _, table := range tables {
		if !migrator.HasTable(table) {
			err := migrator.CreateTable(table)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
