package config

import (
	"WebProject/ginEssential/utils/ContactDb"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func AutoConfig() *sqlx.DB {
	DB := ContactDb.ContactDb()
	fmt.Println("Test Git")
	fmt.Println("321")
	fmt.Println("232")
	return DB
}
