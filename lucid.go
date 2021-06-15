package lucid

import (
	DB "github.com/arthurkay/lucid/db"
)

// DBInstance takes in database initialisation credentials
func DBInstance(user, password, db, host, port, dbType string) *DB.Database {
	return &DB.Database{DBType: dbType, DBUser: user, DBName: db, DBPassword: password, DBHost: host, DBPort: port}
}
