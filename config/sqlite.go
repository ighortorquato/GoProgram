package config

import (
	"os"

	"github.com/ighortorquato/GoProgram.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializrSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	//Check if the database file exist
	_, err := os.Stat("dbPath")
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		//Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Create DB  and connect
	db, err := gorm.Open(sqlite.Open("dbPathb"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqliteopeningerror: %v", err)
		return nil, err
	}
	//Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	//Return the DB
	return db, nil
}
