package helpers

import "gorinha/src/db"

func InitTestDB() {
	SetTestEnv()
	db.Init()
	db.Drop()
	db.Migrate()
}
