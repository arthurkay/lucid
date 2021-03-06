package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database database instance modelling
type Database struct {
	DBType     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBHost     string
}

// MySQLDBConfig initialises a MySQL/MariaDB database instance. Return gorm instance and error
func (db *Database) MySQLDBConfig() (*gorm.DB, error) {
	dbConfig := db.DBUser + ":" + db.DBPassword + "@tcp(" + db.DBHost + ":" + db.DBPort + ")/" + db.DBName + "?parseTime=true"
	connection, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	return connection, err
}

// PostgreSQLConfig initialises a PostgreSQL db instamces, returns gorm, error
func (db *Database) PostgreSQLConfig() (*gorm.DB, error) {
	dbConfig := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s 
                        sslmode=disable`,
		db.DBHost,
		db.DBUser,
		db.DBPassword,
		db.DBName)

	connection, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	return connection, err
}

// DBConfig initialises the required database based on the selected DB from the .env file
func (db *Database) DBConfig() (*gorm.DB, error) {
	switch db.DBType {
	case "mysql":
		return db.MySQLDBConfig()
	case "postgres":
		return db.PostgreSQLConfig()
	default:
		return db.PostgreSQLConfig()
	}
}
