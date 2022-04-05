package apis

import (
	"os"
	"sync"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	// gorm mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type singletonDatabase struct {
	Database *gorm.DB
}

var (
	instanceDatabase *singletonDatabase
	onceDatabase     sync.Once
)

func getDatabase() *gorm.DB {
	return getDbInstance().Database
}

func getDbInstance() *singletonDatabase {
	onceDatabase.Do(func() {

		instanceDatabase = &singletonDatabase{}
		dbString := os.Getenv("DATA_SOURCE_NAME")

		var err error
		instanceDatabase.Database, err = gorm.Open("mysql", dbString)
		if err != nil {
		}

		initDatabase(instanceDatabase.Database)
	})
	return instanceDatabase
}

func initDatabase(db *gorm.DB) {
	db.DB().SetMaxIdleConns(0)
	db.LogMode(true)
}
