package config

import (
	"WebProject/ginEssential/utils/ContactDb"
	"github.com/jmoiron/sqlx"
)

func AutoConfig() *sqlx.DB {
	DB := ContactDb.ContactDb()
	return DB
}
