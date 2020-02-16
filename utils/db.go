package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"sync"
	"time"
)

type dbUtil struct {
	db *gorm.DB
}

var dbInstance *dbUtil
var dbOnce sync.Once

// GetDBConnection gets DB connection
func GetDBConnection() *gorm.DB {
	dbOnce.Do(func() {
		log.Println("Initialize DB connection...")
		conn := os.Getenv("USERNAME_DB") + ":" + os.Getenv("PASSWORD_DB") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME") + "?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(os.Getenv("DATABASE_TYPE"), conn)

		if err != nil {
			panic(err)
		}

		/**
		 * NOTES: this will set connection lifetime in connection pool to 1 minute.
		 * 		  If the connection in the pool is idle > 1 min, Golang will close it
		 * 		  and will create new connection if #connections in the pool < pool max num
		 * 		  of connection. This to avoid invalid connection issue
		 */
		db.DB().SetConnMaxLifetime(time.Second * 60)
		db.SingularTable(true) // Set as singular table
		db.LogMode(false)

		if err != nil {
			panic(err)
		}

		dbInstance = &dbUtil{
			db: db,
		}
	})

	return dbInstance.db
}
